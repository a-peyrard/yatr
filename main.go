package main

import (
	"context"
	"github.com/a-peyrard/yatr/internal/chans"
	"github.com/a-peyrard/yatr/internal/output"
	"github.com/a-peyrard/yatr/internal/worker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var SpecRegex = regexp.MustCompile(`\.spec\.js$`)

func runSpecs(numberOfWorkers int, path string, out output.Output) error {
	start := time.Now()

	ctx, cancelFunc := context.WithCancel(context.Background())

	specChan := make(chan string, numberOfWorkers)
	var (
		wg = sync.WaitGroup{}
	)

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker.Run(ctx, i+1, specChan, &wg, out)
	}

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && SpecRegex.MatchString(path) {
			chans.Send(ctx, specChan, path)
		}
		return nil
	})
	close(specChan)
	wg.Wait()

	cancelFunc()

	out.DisplaySummary(output.Summary{
		SpecPassed: 100,
		SpecFailed: 0,
		TestPassed: 1000,
		TestFailed: 0,
		Elapsed:    time.Since(start),
	})

	return err
}

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	if len(os.Args) < 3 {
		panic("Usage: yatr <numberOfWorkers> <path>")
	}
	numberOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	path := os.Args[2]

	out := output.NewTerminalOutput()
	//out := output.NewNoOpOutput()

	if err := runSpecs(numberOfWorkers, path, out); err != nil {
		panic(err)
	}
}
