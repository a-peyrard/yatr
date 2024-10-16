package worker

import (
	"bufio"
	"context"
	"fmt"
	"github.com/a-peyrard/yatr/internal/output"
	"github.com/rs/zerolog/log"
	"io"
	"os/exec"
	"strings"
	"sync"
)

const RunnerPath = "node-worker/runner-interactive.js"

func Run(ctx context.Context, idx int, workerChan <-chan string, wg *sync.WaitGroup, out output.Output) {
	var worker *exec.Cmd
	defer func() {
		log.Trace().Msgf("Worker %d is finishing... doing cleanup", idx)
		if worker != nil {
			err := worker.Process.Kill()
			if err != nil {
				log.Error().Err(err).Msgf("Failed to kill worker %d: %s", idx, err)
			}
		}
		wg.Done()
	}()

	// create a node worker
	worker = exec.Command("node", RunnerPath)
	workerStdin, err := worker.StdinPipe()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get worker stdin: %s", err)
		return
	}
	workerStdout, err := worker.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get worker stdout: %s", err)
		return
	}
	err = worker.Start()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to start worker: %s", err)
		return
	}

	log.Trace().Msgf("Worker %d ready to do some work...", idx)
	for {
		select {
		case <-ctx.Done():
			log.Trace().Msgf("Worker %d received cancellation signal", idx)
			return
		case spec, ok := <-workerChan:
			if !ok {
				log.Trace().Msgf("Worker %d received close signal", idx)
				return
			}
			err := executeSpec(idx, spec, workerStdin, workerStdout, out)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to execute spec: %s", err)
				return
			}
		}
	}
}

func executeSpec(workerId int, spec string, workerStdin io.WriteCloser, workerStdout io.ReadCloser, out output.Output) error {
	log.Trace().Msgf("Worker %d is running spec: %s", workerId, spec)
	out.StartSpec(workerId, spec)

	scanner := bufio.NewScanner(workerStdout)
	_, err := fmt.Fprintln(workerStdin, spec)
	if err != nil {
		return fmt.Errorf("failed to write to worker stdin: %w", err)
	}

	specOutput := ""
	for scanner.Scan() {
		line := scanner.Text()
		log.Trace().Msgf("Worker %d: RECEIVED %s", workerId, line)
		if strings.HasPrefix(line, "Finished running: ") {
			break
		}
		specOutput += line + "\n"
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read from worker stdout: %w", err)
	}
	log.Trace().Msgf("Worker %d finished running spec: %s", workerId, spec)

	out.FinishSpec(workerId, spec, true)

	return nil
}
