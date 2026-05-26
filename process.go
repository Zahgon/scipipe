package scipipe

// Process is the central component in SciPipe after Workflow. Processes are
// long-running "services" that schedules and executes Tasks based on the IPs
// and parameters received on its in-ports and parameter ports
type Process struct {
	BaseProcess
	CommandPattern string
	PathFuncs      map[string]func(*Task) string
	CustomExecute  func(*Task)
	CoresPerTask   int
	Prepend        string
	Spawn          bool
	PortInfo       map[string]*PortInfo
}

// ------------------------------------------------------------------------
// Factory method(s)
// ------------------------------------------------------------------------

// NewProc returns a new Process, and initializes its ports based on the
// command pattern.
func NewProc(workflow *Workflow, name string, cmd string) *Process {
	_ = "STUB: not implemented"
	return nil
}

// PortInfo is a container for various information about process ports
type PortInfo struct {
	portType  string
	extension string
	doStream  bool
	join      bool
	joinSep   string
}

// initPortsFromCmdPattern is a helper function for NewProc, that sets up in-
// and out-ports based on the shell command pattern used to create the Process.
// Ports are set up in this way:
// `{i:PORTNAME}` specifies an in-port
// `{o:PORTNAME}` specifies an out-port
// `{os:PORTNAME}` specifies an out-port that streams via a FIFO file
// `{p:PORTNAME}` a "parameter (in-)port", which means a port where parameters can be "streamed"
func (p *Process) initPortsFromCmdPattern(cmd string, params map[string]string) {
	_ = "STUB: not implemented"
	// Find in/out port names and params and set up ports
	return
}

// If the |-separated part starts with a dot, treat it as a
// configuration for file extenion to use

// If the |-separated part starts with "join:"
// then treat the character following that as the character to use
// when joining multiple files received on a sub-stream

// initDefaultPathFuncs does exactly what it name says: Initializes default
// path formatters for processes, that is used if no explicit path is set, using
// the proc.SetPath[...] methods
func (p *Process) initDefaultPathFuncs() { _ = "STUB: not implemented"; return }

func sortedFileIPMapKeys(kv map[string]*FileIP) []string { _ = "STUB: not implemented"; return nil }

func sortedStringMapKeys(kv map[string]string) []string { _ = "STUB: not implemented"; return nil }

func sortedFileIPSliceMapKeys(kv map[string][]*FileIP) []string {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------------------
// Main API methods: Port accessor methods
// ------------------------------------------------------------------------

// In is a short-form for InPort() (of BaseProcess), which works only on Process
// processes
func (p *Process) In(portName string) *InPort { _ = "STUB: not implemented"; return nil }

// Return the (only) in-port available

// Out is a short-form for OutPort() (of BaseProcess), which works only on
// Process processes
func (p *Process) Out(portName string) *OutPort { _ = "STUB: not implemented"; return nil }

// Return the (only) out-port available

// InParam is a short-form for InParamPort() (of BaseProcess), which works only on Process
// processes
func (p *Process) InParam(portName string) *InParamPort { _ = "STUB: not implemented"; return nil }

// OutParam is a short-form for OutParamPort() (of BaseProcess), which works only on
// Process processes
func (p *Process) OutParam(portName string) *OutParamPort { _ = "STUB: not implemented"; return nil }

// ------------------------------------------------------------------------
// Main API methods: Configure path formatting
// ------------------------------------------------------------------------

// SetOut initializes a port (if it does not already exist), and takes a
// configuration for its outputs paths via a pattern similar to the command
// pattern used to create new processes, with placeholder tags. Available
// placeholder tags to use are:
// {i:inport_name}
// {p:param_name}
// {t:tag_name}
// An example might be: {i:foo}.replace_with_{p:replacement}.txt
// ... given that the process contains an in-port named 'foo', and a parameter
// named 'replacement'.
// If an out-port with the specified name does not exist, it will be created.
// This allows to create out-ports for filenames that are created without explicitly
// stating a filename on the commandline, such as when only submitting a prefix.
func (p *Process) SetOut(outPortName string, pathPattern string) { _ = "STUB: not implemented"; return }

// Avoiding reusing the same variable in multiple instances of this func

// Replace placeholder with concrete value

// SetOutFunc takes a function which produces a file path based on data
// available in *Task, such as concrete file paths and parameter values,
func (p *Process) SetOutFunc(outPortName string, pathFmtFunc func(task *Task) (path string)) {
	_ = "STUB: not implemented"
	return
}

// ------------------------------------------------------------------------
// Run method
// ------------------------------------------------------------------------

// Run runs the process by instantiating and executing Tasks for all inputs
// and parameter values on its in-ports. in the case when there are no inputs
// or parameter values on the in-ports, it will run just once before it
// terminates. note that the actual execution of shell commands are done inside
// Task.Execute, not here.
func (p *Process) Run() { _ = "STUB: not implemented"; return }

// Check that CoresPerTask is a sane number

// Using a slice to store unprocessed tasks allows us to receive tasks as
// they are produced and to maintain the correct order of IPs. This select
// allows us to process completed tasks as they become available. Waiting
// for all Tasks to be spawned before processing any can cause deadlock
// under certain workflow architectures when there are more than getBufsize()
// Tasks per process, see #81.

// Sending FIFOs for the task

// Execute task in separate go-routine

// Streaming (FIFO) outputs have been sent earlier

// Remove any FIFO file

// createTasks is a helper method for Run that creates tasks based on incoming
// IPs on in-ports, and feeds them to the Run method on the returned channel ch
func (p *Process) createTasks() (ch chan *Task) { _ = "STUB: not implemented"; return nil }

// Tags need to be per Task, otherwise they are overwritten by future IPs

// Only read on in-ports if we have any

// If in-port is closed, that means we got the last params on last iteration, so break

// Only read on param in-ports if we have any

// If param-port is closed, that means we got the last params on last iteration, so break

// Create task and send on the channel we are about to return

// If we have no in-ports nor param in-ports, we should break after the first iteration

type taskQueue []*Task

// NextTaskDone allows us to wait for the next task to be done if it's
// available. Otherwise, nil is returned since nil channels always block.
func (tq taskQueue) NextTaskDone() chan int { _ = "STUB: not implemented"; return nil }
