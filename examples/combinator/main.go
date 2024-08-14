// Workflow written in SciPipe.
// For more information about SciPipe, see: http://scipipe.org
package main

import (
	"fmt"
	"strings"

	"github.com/scipipe/scipipe"
	sp "github.com/scipipe/scipipe"
)

func main() {
	// Create a workflow, using 4 cpu cores
	wf := sp.NewWorkflow("my_workflow", 1)

	comb := NewCombinator(wf, "combinator", func() map[string][]string {
		out := map[string][]string{
			"p:l": []string{},
			"p:n": []string{},
			"p:u": []string{},
		}
		for _, l := range []string{"a", "b", "c"} {
			for _, n := range []string{"1", "2", "3"} {
				for _, u := range []string{"A", "B", "C"} {
					out["p:l"] = append(out["p:l"], l)
					out["p:n"] = append(out["p:n"], n)
					out["p:u"] = append(out["p:u"], u)
				}
			}
		}
		return out
	})

	// Initialize processes
	greeter := wf.NewProc("fooer", "echo {p:l}{p:n}{p:u} > {o:combinations}")
	greeter.InParam("l").From(comb.OutParamPort("l"))
	greeter.InParam("n").From(comb.OutParamPort("n"))
	greeter.InParam("u").From(comb.OutParamPort("u"))
	greeter.SetOutFunc("combinations", func(t *scipipe.Task) string {
		return fmt.Sprintf("%s%s%s.txt", t.Param("l"), t.Param("n"), t.Param("u"))
	})

	// Run the workflow
	wf.Run()
}

type Combinator struct {
	scipipe.BaseProcess
	fun func() map[string][]string
}

func NewCombinator(wf *scipipe.Workflow, name string, newFun func() map[string][]string) *Combinator {
	comb := &Combinator{
		BaseProcess: scipipe.NewBaseProcess(wf, name),
		fun:         newFun,
	}
	for pSpec, _ := range comb.fun() {
		if !strings.Contains(pSpec, ":") {
			scipipe.Fail("You have to specify a type for each output, like f:out_file or p:some_param")
		}
		parts := strings.Split(pSpec, ":")
		pType := map[string]string{
			"f": "file",
			"p": "param",
		}[parts[0]]
		pName := parts[1]
		if pType == "file" {
			if _, ok := comb.OutPorts()[pName]; !ok {
				comb.InitOutPort(comb, pName)
			}
		} else if pType == "param" {
			if _, ok := comb.OutParamPorts()[pName]; !ok {
				comb.InitOutParamPort(comb, pName)
			}
		}
	}
	wf.AddProc(comb)
	return comb
}

func (p *Combinator) Run() {
	defer p.CloseAllOutPorts()

	for pSpec, values := range p.fun() {
		parts := strings.Split(pSpec, ":")
		pType := parts[0]
		pName := parts[1]

		if pType == "f" {
			for _, val := range values {
				ip, err := scipipe.NewFileIP(val)
				scipipe.Check(err)
				p.OutPort(pName).Send(ip)
			}
		} else if pType == "p" {
			for _, val := range values {
				p.OutParamPort(pName).Send(val)
			}
		}
	}
}
