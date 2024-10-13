package main

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var SpecRegex = regexp.MustCompile(`\.spec\.js$`)

func findSpecs(numberOfWorkers int, path string) error {
	ctx, cancelFunc := context.WithCancel(context.Background())

	specChan := make(chan string, numberOfWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, specChan, &wg)
	}

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && SpecRegex.MatchString(path) {
			sendSpec(ctx, specChan, path)
		}
		return nil
	})
	close(specChan)
	wg.Wait()
	cancelFunc()
	return err
}

func worker(ctx context.Context, idx int, workerChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			println("worker", idx, "done")
			return
		case spec, ok := <-workerChan:
			if !ok {
				println("worker", idx, "done")
				return
			}
			time.Sleep(10 * time.Millisecond)
			println("worker", idx, "working on", spec)
		}
	}
}

func sendSpec(ctx context.Context, specChan chan<- string, spec string) {
	select {
	case <-ctx.Done():
		return
	case specChan <- spec:
	}
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: yatr <numberOfWorkers> <path>")
	}
	numberOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	path := os.Args[2]
	if err := findSpecs(numberOfWorkers, path); err != nil {
		panic(err)
	}
}
