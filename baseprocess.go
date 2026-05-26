package scipipe

// BaseProcess provides a skeleton for processes, such as the main Process
// component, and the custom components in the scipipe/components library
type BaseProcess struct {
	name          string
	workflow      *Workflow
	inPorts       map[string]*InPort
	outPorts      map[string]*OutPort
	inParamPorts  map[string]*InParamPort
	outParamPorts map[string]*OutParamPort
}

// NewBaseProcess returns a new BaseProcess, connected to the provided workflow,
// and with the name name
func NewBaseProcess(wf *Workflow, name string) BaseProcess {
	_ = "STUB: not implemented"
	return *new(BaseProcess)
}

// Name returns the name of the process
func (p *BaseProcess) Name() string {
	_ = "STUB: not implemented"

	// Workflow returns the workflow the process is connected to
	return ""
}

func (p *BaseProcess) Workflow() *Workflow {
	_ = "STUB: not implemented"

	// ------------------------------------------------
	// In-port stuff
	// ------------------------------------------------
	return nil
}

// InPort returns the in-port with name portName
func (p *BaseProcess) InPort(portName string) *InPort { _ = "STUB: not implemented"; return nil }

// InitInPort adds the in-port port to the process, with name portName
func (p *BaseProcess) InitInPort(proc WorkflowProcess, portName string) {
	_ = "STUB: not implemented"
	return
}

// InPorts returns a map of all the in-ports of the process, keyed by their
// names
func (p *BaseProcess) InPorts() map[string]*InPort {
	_ = "STUB: not implemented"

	// DeleteInPort deletes an InPort object from the process
	return nil
}

func (p *BaseProcess) DeleteInPort(portName string) { _ = "STUB: not implemented"; return }

// ------------------------------------------------
// Out-port stuff
// ------------------------------------------------

// InitOutPort adds the out-port port to the process, with name portName
func (p *BaseProcess) InitOutPort(proc WorkflowProcess, portName string) {
	_ = "STUB: not implemented"
	return
}

// OutPort returns the out-port with name portName
func (p *BaseProcess) OutPort(portName string) *OutPort { _ = "STUB: not implemented"; return nil }

// OutPorts returns a map of all the out-ports of the process, keyed by their
// names
func (p *BaseProcess) OutPorts() map[string]*OutPort {
	_ = "STUB: not implemented"

	// DeleteOutPort deletes a OutPort object from the process
	return nil
}

func (p *BaseProcess) DeleteOutPort(portName string) { _ = "STUB: not implemented"; return }

// ------------------------------------------------
// Param-in-port stuff
// ------------------------------------------------

// InitInParamPort adds the parameter port paramPort with name portName
func (p *BaseProcess) InitInParamPort(proc WorkflowProcess, portName string) {
	_ = "STUB: not implemented"
	return
}

// InParamPort returns the parameter port with name portName
func (p *BaseProcess) InParamPort(portName string) *InParamPort {
	_ = "STUB: not implemented"
	return nil
}

// InParamPorts returns all parameter in-ports of the process
func (p *BaseProcess) InParamPorts() map[string]*InParamPort { _ = "STUB: not implemented"; return nil }

// DeleteInParamPort deletes a InParamPort object from the process
func (p *BaseProcess) DeleteInParamPort(portName string) { _ = "STUB: not implemented"; return }

// ------------------------------------------------
// Param-out-port stuff
// ------------------------------------------------

// InitOutParamPort initializes the parameter port paramPort with name portName
// to the process We need to supply the concrete process used here as well,
// since this method might be used as part of an embedded struct, meaning that
// the process in the receiver is just the *BaseProcess, which doesn't suffice.
func (p *BaseProcess) InitOutParamPort(proc WorkflowProcess, portName string) {
	_ = "STUB: not implemented"
	return
}

// OutParamPort returns the parameter port with name portName
func (p *BaseProcess) OutParamPort(portName string) *OutParamPort {
	_ = "STUB: not implemented"
	return nil
}

// OutParamPorts returns all parameter out-ports of the process
func (p *BaseProcess) OutParamPorts() map[string]*OutParamPort {
	_ = "STUB: not implemented"
	return nil

	// DeleteOutParamPort deletes a OutParamPort object from the process
}

func (p *BaseProcess) DeleteOutParamPort(portName string) { _ = "STUB: not implemented"; return }

// ------------------------------------------------
// Other stuff
// ------------------------------------------------

// Ready checks whether all the process' ports are connected
func (p *BaseProcess) Ready() (isReady bool) { _ = "STUB: not implemented"; return false }

// CloseOutPorts closes all (normal) out-ports
func (p *BaseProcess) CloseOutPorts() { _ = "STUB: not implemented"; return }

// CloseOutParamPorts closes all parameter out-ports
func (p *BaseProcess) CloseOutParamPorts() { _ = "STUB: not implemented"; return }

// CloseAllOutPorts closes all normal-, and parameter out ports
func (p *BaseProcess) CloseAllOutPorts() { _ = "STUB: not implemented"; return }

// Failf fails with a message that includes the process name
func (p *BaseProcess) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

// Fail fails with a message that includes the process name
func (p *BaseProcess) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

func (p *BaseProcess) Auditf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (p *BaseProcess) Audit(msg interface{}) { _ = "STUB: not implemented"; return }

func (p *BaseProcess) receiveOnInPorts() (ips map[string]*FileIP, inPortsOpen bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Read input IPs on in-ports and set up path mappings

func (p *BaseProcess) receiveOnInParamPorts() (params map[string]string, paramPortsOpen bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Read input IPs on in-ports and set up path mappings
