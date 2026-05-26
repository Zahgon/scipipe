package components

import (
	"github.com/scipipe/scipipe"
)

// ParamSource will feed parameters on an out-port
type ParamSource struct {
	scipipe.BaseProcess
	params []string
}

// NewParamSource returns a new ParamSource
func NewParamSource(wf *scipipe.Workflow, name string, params ...string) *ParamSource {
	_ = "STUB: not implemented"
	return nil
}

// Out returns the out-port, on which parameters the process was initialized
// with, will be retrieved.
func (p *ParamSource) Out() *scipipe.OutParamPort { _ = "STUB: not implemented"; return nil }

// Run runs the process
func (p *ParamSource) Run() { _ = "STUB: not implemented"; return }
