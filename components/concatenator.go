package components

import (
	"github.com/scipipe/scipipe"
)

// Concatenator is a process that concatenates the content of multiple files
// received in the in-port In, into one file returned on its out-port, Out.
// You can optionally specify a tag name to GroupByTag, which will make files
// go into separate output files if they have different values for that tag.
// These output files will have the tag name appended to the base file name.
type Concatenator struct {
	scipipe.BaseProcess
	OutPath    string
	GroupByTag string
}

// NewConcatenator returns a new, initialized Concatenator process
func NewConcatenator(wf *scipipe.Workflow, name string, outPath string) *Concatenator {
	_ = "STUB: not implemented"
	return nil
}

// In returns the (only) in-port for this process
func (p *Concatenator) In() *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// Out returns the (only) out-port for this process
func (p *Concatenator) Out() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the Concatenator process
func (p *Concatenator) Run() { _ = "STUB: not implemented"; return }

// Close file handles

// Send IPs
