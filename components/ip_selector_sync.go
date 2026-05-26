package components

import (
	"github.com/scipipe/scipipe"
	sp "github.com/scipipe/scipipe"
)

// NewIPSelectorSync returns a new IPSelectorSync component.  See the docs for
// IPSelectorSync for more information about how to configure and use it.
func NewIPSelectorSync(wf *sp.Workflow, name string, includeFunc func(ip *sp.FileIP) bool) *IPSelectorSync {
	_ = "STUB: not implemented"
	return nil
}

// IPSelectorSync enables filtering IPs (FileIPs to be specific) by applying the
// supplied function includeFunc, which, if it returns true for an IP, will
// include it.
// The IPSelectorSync requires that the same number and names of ports are used and
// connected both for in-ports and out-ports. So, if you have an in-port
// 'data1', and 'data2', you need to create and connect also out-ports 'data1',
// and 'data2'.
// It will read all in-ports in a synchronous manner, and drop all IPs in the
// current iteration, if the condition in the includeFunc is not met.
type IPSelectorSync struct {
	sp.BaseProcess
	includeFunc func(*sp.FileIP) bool
}

// In returns an in-port if it exists, or creates it before, if it does not exist
func (p *IPSelectorSync) In(name string) *sp.InPort { _ = "STUB: not implemented"; return nil }

// Out returns an out-port if it exists, or creates it before, if it does not exist
func (p *IPSelectorSync) Out(name string) *sp.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the component
func (p *IPSelectorSync) Run() { _ = "STUB: not implemented"; return }

// Send on an out-port with the same name as the in-port

func (p *IPSelectorSync) syncRead() (ipSetChan chan map[string]*scipipe.FileIP) {
	_ = "STUB: not implemented"
	return nil
}

func (p *IPSelectorSync) recvOneEach() (ips map[string]*scipipe.FileIP, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Check if there is any inconsistencies in the OKs (all ports should
// ideally close at the same iteration, otherwise the input streams are not
// in sync).
