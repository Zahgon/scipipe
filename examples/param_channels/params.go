package main

import (
	"runtime"

	sci "github.com/scipipe/scipipe"
)

func main() {
	runtime.GOMAXPROCS(4)
	wf := sci.NewWorkflow("test_wf", 4)

	cmb := NewCombinatoricsGen(wf, "combgen")

	// An abc file printer
	abc := wf.NewProc("abc", "echo {p:a} {p:b} {p:c} > {o:out}; sleep 1")
	abc.Spawn = true
	abc.SetOut("out", "{p:a}_{p:b}_{p:c}.txt")

	// A printer task
	prt := wf.NewProc("printer", "cat {i:in} >> log.txt")
	prt.Spawn = false

	// Connection info
	abc.InParamPort("a").From(cmb.A())
	abc.InParamPort("b").From(cmb.B())
	abc.InParamPort("c").From(cmb.C())
	prt.In("in").From(abc.Out("out"))

	wf.Run()
}

type CombinatoricsGen struct {
	sci.BaseProcess
}

func NewCombinatoricsGen(wf *sci.Workflow, name string) *CombinatoricsGen {
	_ = "STUB: not implemented"
	return nil
}

func (p *CombinatoricsGen) A() *sci.OutParamPort { _ = "STUB: not implemented"; return nil }
func (p *CombinatoricsGen) B() *sci.OutParamPort { _ = "STUB: not implemented"; return nil }
func (p *CombinatoricsGen) C() *sci.OutParamPort { _ = "STUB: not implemented"; return nil }

func (p *CombinatoricsGen) Run() { _ = "STUB: not implemented"; return }
