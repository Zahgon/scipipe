package main

import (
	sci "github.com/scipipe/scipipe"
)

func main() {
	// Initiate
	wfl := sci.NewWorkflow("wrapperwf", 4)
	foo := NewFooer(wfl, "fooer")
	f2b := NewFoo2Barer(wfl, "foo2barer")

	// Connect
	f2b.InFoo().From(foo.OutFoo())

	// Run
	wfl.Run()
}

// ------------------------------------------------
// Components
// ------------------------------------------------

// Fooer
// -----

type Fooer struct {
	*sci.Process
	name string
}

func NewFooer(wf *sci.Workflow, name string) *Fooer { _ = "STUB: not implemented"; return nil }

// Define static ports

func (p *Fooer) OutFoo() *sci.OutPort { _ = "STUB: not implemented"; return nil }

// Foo2Barer
// ---------

type Foo2Barer struct {
	*sci.Process
	name string
}

func NewFoo2Barer(wf *sci.Workflow, name string) *Foo2Barer { _ = "STUB: not implemented"; return nil }

// Define static ports

func (p *Foo2Barer) InFoo() *sci.InPort { _ = "STUB: not implemented"; return nil }

func (p *Foo2Barer) OutBar() *sci.OutPort { _ = "STUB: not implemented"; return nil }
