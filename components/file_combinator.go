package components

import (
	"github.com/scipipe/scipipe"
)

// FileCombinator takes a set of input streams of FileIPs, and returns the same
// number of output streams, where the FileIPs are multiplied so as to
// guarantee that all combinations of the ips in the input streams are created.
// Input ports and corresponding out-ports (with the same port names) are
// created on demand, by accessing them with the p.In(PORTNAME) method.
// The corresponding out-porta can then be accessed with the same port name
// with p.Out(PORTNAME)
type FileCombinator struct {
	scipipe.BaseProcess
	globPatterns []string
}

// NewFileCombinator returns a new initialized FileCombinator process
func NewFileCombinator(wf *scipipe.Workflow, name string) *FileCombinator {
	_ = "STUB: not implemented"
	return nil
}

// In returns the in-port with name pName. If it does not exist, it will create
// that in-port, and a corresponding out-port with the same port name.
func (p *FileCombinator) In(pName string) *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// Initialize a corresponding outport with the same name, for each inport

// Out returns the outport
func (p *FileCombinator) Out(pName string) *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the FileCombinator process
func (p *FileCombinator) Run() { _ = "STUB: not implemented"; return }

// Collect all the input IPs

// Send combinations of all IPs

// Make unique copy of variables for this iteration, so they don't get
// overwritten on the next loop iteration

// combine is a recursive method that creates combinations of all the IPs in the input IP arrays, such that:
// [a.txt b.txt]
// [1,txt 2.txt 3.txt]
// ... will be turned into:
// [a.txt a.txt a.txt b.txt b.txt b.txt]
// [1.txt 2.txt 3.txt 1.txt 2.txt 3.txt]
// as an example.
func (p *FileCombinator) combine(inIPs map[string][]*scipipe.FileIP, keys []string) map[string][]*scipipe.FileIP {
	_ = "STUB: not implemented"
	return nil
}

// Recursive call

// Multiply each string in head with the length of the rows in the tail
// (they are guaranteed to be of equal length)

// Multiply the content of each row in the tail with the number of rows
// in the tail
