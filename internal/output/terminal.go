package output

import (
	"fmt"
	"path"
	"strings"
	"sync"
	"time"
)

const (
	Reset           = "\033[0m"
	Bold            = "\033[1m"
	Dim             = "\033[2m"
	UnderlineOff    = "\033[24m"
	ReverseVideoOn  = "\033[7m"
	ReverseVideoOff = "\033[27m"

	YellowText   = "\033[33m"
	GreenText    = "\033[32m"
	RedText      = "\033[31m"
	ClearLine    = "\033[K"
	MoveCursorUp = "\033[%dA"
)

type (
	Summary struct {
		SpecPassed int
		SpecFailed int
		TestPassed int
		TestFailed int
		Elapsed    time.Duration
	}
	Output interface {
		StartSpec(workerId int, specName string)
		FinishSpec(workerId int, specName string, success bool)
		DisplaySummary(summary Summary)
	}
	TerminalOutput struct {
		inProgress map[string]struct{}
		mu         sync.Mutex
	}
	NoOpOutput struct{}
)

func (n NoOpOutput) StartSpec(workerId int, specName string) {}

func (n NoOpOutput) FinishSpec(workerId int, specName string, success bool) {}

func (n NoOpOutput) DisplaySummary(summary Summary) {}

func NewTerminalOutput() Output {
	return &TerminalOutput{
		inProgress: make(map[string]struct{}),
	}
}

func NewNoOpOutput() Output {
	return NoOpOutput{}
}

// StartSpec sets the worker to be running a spec
func (o *TerminalOutput) StartSpec(workerId int, specName string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if _, ok := o.inProgress[specName]; ok {
		return
	}
	o.inProgress[specName] = struct{}{}

	// update the display
	printLine(workerId, "runs", YellowText, specName)
}

// FinishSpec updates the worker status when the spec finishes
func (o *TerminalOutput) FinishSpec(workerId int, specName string, success bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	delete(o.inProgress, specName)

	// update the display
	fmt.Printf(MoveCursorUp, len(o.inProgress)+1)
	if success {
		printLine(workerId, "pass", GreenText, specName)
	} else {
		printLine(workerId, "fail", RedText, specName)
	}
	for specName := range o.inProgress {
		printLine(workerId, "runs", YellowText, specName)
	}
}

func (o *TerminalOutput) DisplaySummary(s Summary) {
	fmt.Printf("\n%sTest Suites: %s%s%d passed%s, %d total%s\n", Bold, GreenText, Bold, s.SpecPassed, Reset, s.SpecPassed+s.SpecFailed, Reset)
	fmt.Printf("%sTests:       %s%s%d passed%s, %d total%s\n", Bold, GreenText, Bold, s.TestPassed, Reset, s.TestPassed+s.TestFailed, Reset)
	fmt.Printf("%sTime:        %s%s\n", Bold, Reset, s.Elapsed)
	fmt.Printf("%sRan all test suites%s.\n", Dim, Reset)
}

func printLine(workerId int, status string, color string, specName string) {
	dir := path.Dir(specName)
	fileName := path.Base(specName)

	fmt.Printf(
		"%s%s%s%s %s %s%s %s%s%s%s\n",
		Reset,
		ReverseVideoOn,
		color,
		Bold,
		strings.ToUpper(status),
		Reset,
		Dim,
		dir+"/",
		Bold,
		fileName,
		Reset,
	)
}
