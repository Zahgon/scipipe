package main

import (
	"fmt"

	sci "github.com/scipipe/scipipe"
)

func main() {
	p := NewFooToBarReplacer()
	fmt.Println("Process: ", p)
}

// -------------------------------------------
//  Example of defining a new wrapper task
// -------------------------------------------

type FooToBarReplacer struct {
	InnerProcess *sci.WorkflowProcess
	Run          func(p *FooToBarReplacer)
	InFoo        chan *sci.FileIP
	OutBar       chan *sci.FileIP
}

func NewFooToBarReplacer() interface{} { _ = "STUB: not implemented"; return nil }

// -------------------------------------------
//  New helper methods
// -------------------------------------------

func NewProcessFromStruct(procStruct interface{}, execFunc func(*sci.Task), pathFuncs map[string]func(*sci.Task) string) interface{} {
	_ = "STUB: not implemented"
	// Get in-ports of struct
	return nil
}

// TODO: Change this!

// TODO: Change this!
