package main

import (
	. "github.com/scipipe/scipipe"
)

func main() {
	wf := NewWorkflow("FuncHookWf", 4)

	foo := NewFooer(wf, "foo")
	f2b := NewFoo2Barer(wf, "f2b")

	foo.OutFoo().To(f2b.InFoo())

	wf.Run()
}

// ------------------------------------------------------------------------
// Components
// ------------------------------------------------------------------------

// Fooer

type Fooer struct {
	*Process
	name string
}

func NewFooer(wf *Workflow, name string) *Fooer {
	_ = "STUB: not implemented"
	// Initiate task from a "shell like" pattern, though here we
	// just specify the out-port foo
	return nil
}

// Set the output formatter to a static string

// Create the custom execute function, with pure Go code

// Connect the ports of the outer task to the inner, generic one

func (p *Fooer) OutFoo() *OutPort {
	_ = "STUB: not implemented"

	// Foo2Barer
	return nil
}

type Foo2Barer struct {
	*Process
	name string
}

func NewFoo2Barer(wf *Workflow, name string) *Foo2Barer {
	_ = "STUB: not implemented"
	// Initiate task from a "shell like" pattern, though here we
	// just specify the in-port foo and the out-port bar
	return nil
}

// Set the output formatter to extend the path on the "bar"" in-port

// Create the custom execute function, with pure Go code

// Connect the ports of the outer task to the inner, generic one

func (p *Foo2Barer) InFoo() *InPort   { _ = "STUB: not implemented"; return nil }
func (p *Foo2Barer) OutBar() *OutPort { _ = "STUB: not implemented"; return nil }
