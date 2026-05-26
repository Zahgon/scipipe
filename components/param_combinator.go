package components

import (
	"github.com/scipipe/scipipe"
)

// ParamCombinator takes a set of input params, and returns the same
// number of param streams, where the params are multiplied so as to
// guarantee that all combinations of the params in the streams are created.
// Input ports and corresponding out-ports (with the same port names) are
// created on demand, by accessing them with the p.InParam(PORTNAME) method.
// The corresponding out-porta can then be accessed with the same port name
// with p.OutParam(PORTNAME)
type ParamCombinator struct {
	scipipe.BaseProcess
}

// NewParamCombinator returns a new initialized ParamCombinator process
func NewParamCombinator(wf *scipipe.Workflow, name string) *ParamCombinator {
	_ = "STUB: not implemented"
	return nil
}

// InParam returns the in-port with name pName. If it does not exist, it will create
// that in-port, and a corresponding out-port with the same port name.
func (p *ParamCombinator) InParam(pName string) *scipipe.InParamPort {
	_ = "STUB: not implemented"
	return nil
}

// Initialize a corresponding outport with the same name, for each inport

// OutParam returns the outport
func (p *ParamCombinator) OutParam(pName string) *scipipe.OutParamPort {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the ParamCombinator process
func (p *ParamCombinator) Run() { _ = "STUB: not implemented"; return }

// Collect all input params

// Send combinations of all IPs

// Make unique copy of variables for this iteration, so they don't get
// overwritten on the next loop iteration

// combine is a recursive method that creates combinations of all the IPs in the input IP arrays, such that:
// [a b]
// [1 2 3]
// ... will be turned into:
// [a a a b b b]
// [1 2 3 1 2 3]
// as an example.
func combine(inParams map[string][]string, keys []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// Recursive call

// Multiply each string in head with the length of the rows in the tail
// (they are guaranteed to be of equal length)

// Multiply the content of each row in the tail with the number of rows
// in the tail
