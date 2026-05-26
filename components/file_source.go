package components

import (
	"github.com/scipipe/scipipe"
)

// FileSource is initiated with a set of file paths, which it will send as a
// stream of File IPs on its outport Out()
type FileSource struct {
	scipipe.BaseProcess
	filePaths []string
}

// NewFileSource returns a new initialized FileSource process
func NewFileSource(wf *scipipe.Workflow, name string, filePaths ...string) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

// Out returns the out-port, on which file IPs based on the file paths the
// process was initialized with, will be retrieved.
func (p *FileSource) Out() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the FileSource process
func (p *FileSource) Run() { _ = "STUB: not implemented"; return }
