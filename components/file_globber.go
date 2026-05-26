package components

import (
	"github.com/scipipe/scipipe"
)

// FileGlobber is initiated with a set of glob patterns paths, which it will
// use to find concrete file paths, for which it will return a stream of
// corresponding File IPs on its outport Out()
type FileGlobber struct {
	scipipe.BaseProcess
	globPatterns []string
}

// NewFileGlobber returns a new initialized FileGlobber process
func NewFileGlobber(wf *scipipe.Workflow, name string, globPatterns ...string) *FileGlobber {
	_ = "STUB: not implemented"
	return nil
}

// NewFileGlobberDependent returns a new FileGlobber that depends on upstream
// files to be received on the InPort InDependency() before it starts globbing files.
func NewFileGlobberDependent(wf *scipipe.Workflow, name string, globPatterns ...string) *FileGlobber {
	_ = "STUB: not implemented"
	return nil
}

// Out returns the out-port, on which file IPs based on the file paths the
// process was initialized with, will be retrieved.
func (p *FileGlobber) Out() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// InDependency takes files which it will wait for before it starts to execute.
func (p *FileGlobber) InDependency() *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// Run runs the FileGlobber process
func (p *FileGlobber) Run() { _ = "STUB: not implemented"; return }

// If we have an InDependency in-port, then loop on the in-channel of that, to make
// the process wait for IPs on that.

// Do nothing, just empty the channel

func (p *FileGlobber) globFiles() { _ = "STUB: not implemented"; return }
