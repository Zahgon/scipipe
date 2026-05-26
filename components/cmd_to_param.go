package components

import (
	"github.com/scipipe/scipipe"
)

// CommandToParams takes a shell command, runs it, and sens each of its files
// as parameters on its OutParam parameter port
type CommandToParams struct {
	scipipe.BaseProcess
	command string
}

// NewCommandToParams returns an initialized new CommandToParams
func NewCommandToParams(wf *scipipe.Workflow, name string, command string) *CommandToParams {
	_ = "STUB: not implemented"
	return nil
}

// OutParam returns an parameter out-port with lines of the files being read
func (p *CommandToParams) OutParam() *scipipe.OutParamPort { _ = "STUB: not implemented"; return nil }

// Run the CommandToParams
func (p *CommandToParams) Run() { _ = "STUB: not implemented"; return }
