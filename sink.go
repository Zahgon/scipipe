package scipipe

// Sink is a simple component that just receives IPs on its In-port without
// doing anything with them. It is used to drive pipelines of processes
type Sink struct {
	BaseProcess
}

// NewSink returns a new Sink component
func NewSink(wf *Workflow, name string) *Sink { _ = "STUB: not implemented"; return nil }

func (p *Sink) in() *InPort           { _ = "STUB: not implemented"; return nil }
func (p *Sink) paramIn() *InParamPort { _ = "STUB: not implemented"; return nil }

// From connects an out-port to the sinks in-port
func (p *Sink) From(outPort *OutPort) { _ = "STUB: not implemented"; return }

// FromParam connects a param-out-port to the sinks param-in-port
func (p *Sink) FromParam(outParamPort *OutParamPort) { _ = "STUB: not implemented"; return }

// Run runs the Sink process
func (p *Sink) Run() { _ = "STUB: not implemented"; return }
