// An example that shows how to create a sub-network / sub-workflow that can be
// used as a component
package main

import sp "github.com/scipipe/scipipe"

func main() {
	// Main workflow
	wfl := sp.NewWorkflow("foobar_wf", 4)

	// Sub-workflow
	NewFooBarSubWorkflow(wfl, "foobar_subwf")

	// Run
	wfl.Run()
}

// ------------------------------------------------
// FooBarSubWorkflow
// ------------------------------------------------

type FooBarSubWorkflow struct {
	name  string
	Procs map[string]*sp.Process
	Out   *sp.OutPort
}

func NewFooBarSubWorkflow(wf *sp.Workflow, name string) *FooBarSubWorkflow {
	_ = "STUB: not implemented"
	return nil
}

// Connect together inner processes

// Connect last port of inner process to subnetwork out-port

func (wf *FooBarSubWorkflow) Name() string { _ = "STUB: not implemented"; return "" }

func (wf *FooBarSubWorkflow) Run() { _ = "STUB: not implemented"; return }

func (wf *FooBarSubWorkflow) Ready() bool { _ = "STUB: not implemented"; return false }
