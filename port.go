package scipipe

import (
	"sync"
)

// ------------------------------------------------------------------------
// InPort
// ------------------------------------------------------------------------

// InPort represents a pluggable connection to multiple out-ports from other
// processes, from its own process, and with which it is communicating via
// channels under the hood
type InPort struct {
	Chan        chan *FileIP
	name        string
	process     WorkflowProcess
	RemotePorts map[string]*OutPort
	ready       bool
	closeLock   sync.Mutex
}

// NewInPort returns a new InPort struct
func NewInPort(name string) *InPort { _ = "STUB: not implemented"; return nil }

// This one will contain merged inputs from inChans

// Name returns the name of the InPort
func (pt *InPort) Name() string { _ = "STUB: not implemented"; return "" }

// Process returns the process connected to the port
func (pt *InPort) Process() WorkflowProcess {
	_ = "STUB: not implemented"
	return *new(WorkflowProcess)
}

// SetProcess sets the process of the port to p
func (pt *InPort) SetProcess(p WorkflowProcess) {
	_ = "STUB: not implemented"

	// AddRemotePort adds a remote OutPort to the InPort
	return
}

func (pt *InPort) AddRemotePort(rpt *OutPort) { _ = "STUB: not implemented"; return }

// From connects an OutPort to the InPort
func (pt *InPort) From(rpt *OutPort) { _ = "STUB: not implemented"; return }

// Disconnect disconnects the (out-)port with name rptName, from the InPort
func (pt *InPort) Disconnect(rptName string) { _ = "STUB: not implemented"; return }

// removeRemotePort removes the (out-)port with name rptName, from the InPort
func (pt *InPort) removeRemotePort(rptName string) { _ = "STUB: not implemented"; return }

// SetReady sets the ready status of the InPort
func (pt *InPort) SetReady(ready bool) {
	_ = "STUB: not implemented"

	// Ready tells whether the port is ready or not
	return
}

func (pt *InPort) Ready() bool {
	_ = "STUB: not implemented"

	// Send sends IPs to the in-port, and is supposed to be called from the remote
	// (out-) port, to send to this in-port
	return false
}

func (pt *InPort) Send(ip *FileIP) {
	_ = "STUB: not implemented"

	// Recv receives IPs from the port
	return
}

func (pt *InPort) Recv() *FileIP {
	_ = "STUB: not implemented"

	// CloseConnection closes the connection to the remote out-port with name
	// rptName, on the InPort
	return nil
}

func (pt *InPort) CloseConnection(rptName string) { _ = "STUB: not implemented"; return }

// Failf fails with a message that includes the process name
func (pt *InPort) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

// Fail fails with a message that includes the process name
func (pt *InPort) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// OutPort
// ------------------------------------------------------------------------

// OutPort represents a pluggable connection to multiple in-ports from other
// processes, from its own process, and with which it is communicating via
// channels under the hood
type OutPort struct {
	name        string
	process     WorkflowProcess
	RemotePorts map[string]*InPort
	ready       bool
}

// NewOutPort returns a new OutPort struct
func NewOutPort(name string) *OutPort { _ = "STUB: not implemented"; return nil }

// Name returns the name of the OutPort
func (pt *OutPort) Name() string { _ = "STUB: not implemented"; return "" }

// Process returns the process connected to the port
func (pt *OutPort) Process() WorkflowProcess {
	_ = "STUB: not implemented"
	return *new(WorkflowProcess)
}

// SetProcess sets the process of the port to p
func (pt *OutPort) SetProcess(p WorkflowProcess) {
	_ = "STUB: not implemented"

	// AddRemotePort adds a remote InPort to the OutPort
	return
}

func (pt *OutPort) AddRemotePort(rpt *InPort) { _ = "STUB: not implemented"; return }

// removeRemotePort removes the (in-)port with name rptName, from the OutPort
func (pt *OutPort) removeRemotePort(rptName string) { _ = "STUB: not implemented"; return }

// To connects an InPort to the OutPort
func (pt *OutPort) To(rpt *InPort) { _ = "STUB: not implemented"; return }

// Disconnect disconnects the (in-)port with name rptName, from the OutPort
func (pt *OutPort) Disconnect(rptName string) { _ = "STUB: not implemented"; return }

// SetReady sets the ready status of the OutPort
func (pt *OutPort) SetReady(ready bool) {
	_ = "STUB: not implemented"

	// Ready tells whether the port is ready or not
	return
}

func (pt *OutPort) Ready() bool {
	_ = "STUB: not implemented"

	// Send sends an FileIP to all the in-ports connected to the OutPort
	return false
}

func (pt *OutPort) Send(ip *FileIP) { _ = "STUB: not implemented"; return }

// Close closes the connection between this port and all the ports it is
// connected to. If this port is the last connected port to an in-port, that
// in-ports channel will also be closed.
func (pt *OutPort) Close() { _ = "STUB: not implemented"; return }

// Failf fails with a message that includes the process name
func (pt *OutPort) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

// Fail fails with a message that includes the process name
func (pt *OutPort) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// InParamPort
// ------------------------------------------------------------------------

// InParamPort is an in-port for parameter values of string type
type InParamPort struct {
	Chan        chan string
	name        string
	process     WorkflowProcess
	RemotePorts map[string]*OutParamPort
	ready       bool
	closeLock   sync.Mutex
}

// NewInParamPort returns a new InParamPort
func NewInParamPort(name string) *InParamPort { _ = "STUB: not implemented"; return nil }

// Name returns the name of the InParamPort
func (pip *InParamPort) Name() string { _ = "STUB: not implemented"; return "" }

// Process returns the process that is connected to the port
func (pip *InParamPort) Process() WorkflowProcess {
	_ = "STUB: not implemented"
	return *new(WorkflowProcess)
}

// SetProcess sets the process of the port to p
func (pip *InParamPort) SetProcess(p WorkflowProcess) {
	_ = "STUB: not implemented"

	// AddRemotePort adds a remote OutParamPort to the InParamPort
	return
}

func (pip *InParamPort) AddRemotePort(pop *OutParamPort) { _ = "STUB: not implemented"; return }

// From connects one parameter port with another one
func (pip *InParamPort) From(pop *OutParamPort) { _ = "STUB: not implemented"; return }

// FromStr feeds one or more parameters of type string to a port
func (pip *InParamPort) FromStr(strings ...string) { _ = "STUB: not implemented"; return }

// FromInt feeds one or more parameters of type int to the param port
func (pip *InParamPort) FromInt(ints ...int) { _ = "STUB: not implemented"; return }

// FromFloat feeds one or more parameters of type float64 to the param port
func (pip *InParamPort) FromFloat(floats ...float64) { _ = "STUB: not implemented"; return }

// SetReady sets the ready status of the InParamPort
func (pip *InParamPort) SetReady(ready bool) {
	_ = "STUB: not implemented"

	// Ready tells whether the port is ready or not
	return
}

func (pip *InParamPort) Ready() bool {
	_ = "STUB: not implemented"

	// Send sends IPs to the in-port, and is supposed to be called from the remote
	// (out-) port, to send to this in-port
	return false
}

func (pip *InParamPort) Send(param string) {
	_ = "STUB: not implemented"

	// Recv receiveds a param value over the ports connection
	return
}

func (pip *InParamPort) Recv() string {
	_ = "STUB: not implemented"

	// CloseConnection closes the connection to the remote out-port with name
	// popName, on the InParamPort
	return ""
}

func (pip *InParamPort) CloseConnection(popName string) { _ = "STUB: not implemented"; return }

// Failf fails with a message that includes the process name
func (pt *InParamPort) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

// Fail fails with a message that includes the process name
func (pt *InParamPort) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// OutParamPort
// ------------------------------------------------------------------------

// OutParamPort is an out-port for parameter values of string type
type OutParamPort struct {
	name        string
	process     WorkflowProcess
	RemotePorts map[string]*InParamPort
	ready       bool
}

// NewOutParamPort returns a new OutParamPort
func NewOutParamPort(name string) *OutParamPort { _ = "STUB: not implemented"; return nil }

// Name returns the name of the OutParamPort
func (pop *OutParamPort) Name() string { _ = "STUB: not implemented"; return "" }

// Process returns the process that is connected to the port
func (pop *OutParamPort) Process() WorkflowProcess {
	_ = "STUB: not implemented"
	return *new(WorkflowProcess)
}

// SetProcess sets the process of the port to p
func (pop *OutParamPort) SetProcess(p WorkflowProcess) {
	_ = "STUB: not implemented"

	// AddRemotePort adds a remote InParamPort to the OutParamPort
	return
}

func (pop *OutParamPort) AddRemotePort(pip *InParamPort) { _ = "STUB: not implemented"; return }

// To connects an InParamPort to the OutParamPort
func (pop *OutParamPort) To(pip *InParamPort) { _ = "STUB: not implemented"; return }

// Disconnect disonnects the (in-)port with name rptName, from the OutParamPort
func (pop *OutParamPort) Disconnect(pipName string) { _ = "STUB: not implemented"; return }

// removeRemotePort removes the (in-)port with name rptName, from the OutParamPort
func (pop *OutParamPort) removeRemotePort(pipName string) { _ = "STUB: not implemented"; return }

// SetReady sets the ready status of the OutParamPort
func (pop *OutParamPort) SetReady(ready bool) {
	_ = "STUB: not implemented"

	// Ready tells whether the port is ready or not
	return
}

func (pop *OutParamPort) Ready() bool {
	_ = "STUB: not implemented"

	// Send sends an FileIP to all the in-ports connected to the OutParamPort
	return false
}

func (pop *OutParamPort) Send(param string) { _ = "STUB: not implemented"; return }

// Close closes the connection between this port and all the ports it is
// connected to. If this port is the last connected port to an in-port, that
// in-ports channel will also be closed.
func (pop *OutParamPort) Close() { _ = "STUB: not implemented"; return }

// Failf fails with a message that includes the process name
func (pt *OutParamPort) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

// Fail fails with a message that includes the process name
func (pt *OutParamPort) Fail(msg interface{}) { _ = "STUB: not implemented"; return }
