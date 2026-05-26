package components

import (
	"github.com/scipipe/scipipe"
)

// StreamToSubStream takes a normal stream of IP's representing
// individual files, and returns one IP where the incoming IPs
// are sent on its substream.
type StreamToSubStream struct {
	scipipe.BaseProcess
}

// NewStreamToSubStream instantiates a new StreamToSubStream process
func NewStreamToSubStream(wf *scipipe.Workflow, name string) *StreamToSubStream {
	_ = "STUB: not implemented"
	return nil
}

// In returns the in-port
func (p *StreamToSubStream) In() *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// OutSubStream returns the out-port
func (p *StreamToSubStream) OutSubStream() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the StreamToSubStream
func (p *StreamToSubStream) Run() { _ = "STUB: not implemented"; return }

// create a temporary file, with a _scipipe prefix
