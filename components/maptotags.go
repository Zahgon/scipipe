package components

import (
	"github.com/scipipe/scipipe"
)

// MapToTags is a process that runs a function provided by the user, upon
// initialization, that will provide a map of tag:value pairs, based in IPs read
// on the In-port. The tag:value pairs (maps) are added to the IPs on the
// out-port, which are identical to the incoming IPs, except for the new
// tag:value map
type MapToTags struct {
	scipipe.BaseProcess
	mapFunc func(ip *scipipe.FileIP) map[string]string
}

// NewMapToTags returns an initialized MapToTags process
func NewMapToTags(wf *scipipe.Workflow, name string, mapFunc func(ip *scipipe.FileIP) map[string]string) *MapToTags {
	_ = "STUB: not implemented"
	return nil
}

// In takes input files the content of which the map function will be run,
// to generate tags
func (p *MapToTags) In() *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// Out outputs files that are supplemented with tags by the map function.
func (p *MapToTags) Out() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the MapToTags process
func (p *MapToTags) Run() { _ = "STUB: not implemented"; return }
