// Package scipipe is a library for writing scientific workflows (sometimes
// also called "pipelines") of shell commands that depend on each other, in the
// Go programming languages. It was initially designed for problems in
// cheminformatics and bioinformatics, but should apply equally well to any
// domain involving complex pipelines of interdependent shell commands.
package scipipe

import (
	"sync"
)

// ----------------------------------------------------------------------------
// Workflow
// ----------------------------------------------------------------------------

// Workflow is the centerpiece of the functionality in SciPipe, and is a
// container for a pipeline of processes making up a workflow. It has various
// methods for coordination the execution of the pipeline as a whole, such as
// keeping track of the maxiumum number of concurrent tasks, as well as helper
// methods for creating new processes, that automatically gets plugged in to the
// workflow on creation
type Workflow struct {
	name              string
	procs             map[string]WorkflowProcess
	concurrentTasks   chan struct{}
	concurrentTasksMx sync.Mutex
	sink              *Sink
	driver            WorkflowProcess
	logFile           string
	PlotConf          WorkflowPlotConf
}

// WorkflowPlotConf contains configuraiton for plotting the workflow as a graph
// with graphviz
type WorkflowPlotConf struct {
	EdgeLabels bool
}

// WorkflowProcess is an interface for processes to be handled by Workflow
type WorkflowProcess interface {
	Name() string
	InPorts() map[string]*InPort
	OutPorts() map[string]*OutPort
	InParamPorts() map[string]*InParamPort
	OutParamPorts() map[string]*OutParamPort
	Ready() bool
	Run()
	Fail(interface{})
	Failf(string, ...interface{})
}

// ----------------------------------------------------------------------------
// Factory function(s)
// ----------------------------------------------------------------------------

// NewWorkflow returns a new Workflow
func NewWorkflow(name string, maxConcurrentTasks int) *Workflow {
	_ = "STUB: not implemented"
	return nil
}

// Set up logging

// NewWorkflowCustomLogFile returns a new Workflow, with
func NewWorkflowCustomLogFile(name string, maxConcurrentTasks int, logFile string) *Workflow {
	_ = "STUB: not implemented"
	return nil
}

func newWorkflowWithoutLogging(name string, maxConcurrentTasks int) *Workflow {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Main API methods
// ----------------------------------------------------------------------------

// Name returns the name of the workflow
func (wf *Workflow) Name() string {
	_ = "STUB: not implemented"

	// NewProc returns a new process based on a commandPattern (See the
	// documentation for scipipe.NewProcess for more details about the pattern) and
	// connects the process to the workflow
	return ""
}

func (wf *Workflow) NewProc(procName string, commandPattern string) *Process {
	_ = "STUB: not implemented"
	return nil
}

// Proc returns the process with name procName from the workflow
func (wf *Workflow) Proc(procName string) WorkflowProcess {
	_ = "STUB: not implemented"
	return *new(WorkflowProcess)
}

// ProcsSorted returns the processes of the workflow, in an array, sorted by the
// process names
func (wf *Workflow) ProcsSorted() []WorkflowProcess { _ = "STUB: not implemented"; return nil }

// Procs returns a map of all processes keyed by their names in the workflow
func (wf *Workflow) Procs() map[string]WorkflowProcess {
	_ = "STUB: not implemented"

	// AddProc adds a Process to the workflow, to be run when the workflow runs
	return nil
}

func (wf *Workflow) AddProc(proc WorkflowProcess) { _ = "STUB: not implemented"; return }

// AddProcs takes one or many Processes and adds them to the workflow, to be run
// when the workflow runs.
func (wf *Workflow) AddProcs(procs ...WorkflowProcess) { _ = "STUB: not implemented"; return }

// Sink returns the sink process of the workflow
func (wf *Workflow) Sink() *Sink {
	_ = "STUB: not implemented"

	// SetSink sets the sink of the workflow to the provided sink process
	return nil
}

func (wf *Workflow) SetSink(sink *Sink) { _ = "STUB: not implemented"; return }

// IncConcurrentTasks increases the conter for how many concurrent tasks are
// currently running in the workflow
func (wf *Workflow) IncConcurrentTasks(slots int) {
	_ = "STUB: not implemented"
	// We must lock so that multiple processes don't end up with partially "filled slots"
	return
}

// DecConcurrentTasks decreases the conter for how many concurrent tasks are
// currently running in the workflow
func (wf *Workflow) DecConcurrentTasks(slots int) { _ = "STUB: not implemented"; return }

// PlotGraph writes the workflow structure to a dot file
func (wf *Workflow) PlotGraph(filePath string) { _ = "STUB: not implemented"; return }

// PlotGraphPDF writes the workflow structure to a dot file, and also runs the
// graphviz dot command to produce a PDF file (requires graphviz, with the dot
// command, installed on the system)
func (wf *Workflow) PlotGraphPDF(filePath string) { _ = "STUB: not implemented"; return }

// DotGraph generates a graph description in DOT format
// (See https://en.wikipedia.org/wiki/DOT_%28graph_description_language%29)
// If Workflow.PlotConf.EdgeLabels is set to true, a label containing the
// in-port and out-port to which edges are connected to, will be printed.
func (wf *Workflow) DotGraph() (dot string) { _ = "STUB: not implemented"; return "" }

// File connections

// Parameter connections

// ----------------------------------------------------------------------------
// Run methods
// ----------------------------------------------------------------------------

// Run runs all the processes of the workflow
func (wf *Workflow) Run() { _ = "STUB: not implemented"; return }

// RunTo runs all processes upstream of, and including, the process with
// names provided as arguments
func (wf *Workflow) RunTo(finalProcNames ...string) { _ = "STUB: not implemented"; return }

// RunToRegex runs all processes upstream of, and including, the process
// whose name matches any of the provided regexp patterns
func (wf *Workflow) RunToRegex(procNamePatterns ...string) { _ = "STUB: not implemented"; return }

// RunToProcs runs all processes upstream of, and including, the process strucs
// provided as arguments
func (wf *Workflow) RunToProcs(finalProcs ...WorkflowProcess) { _ = "STUB: not implemented"; return }

// ----------------------------------------------------------------------------
// Helper methods for running the workflow
// ----------------------------------------------------------------------------

// runProcs runs a specified set of processes only
func (wf *Workflow) runProcs(procs map[string]WorkflowProcess) { _ = "STUB: not implemented"; return }

func (wf *Workflow) readyToRun(procs map[string]WorkflowProcess) bool {
	_ = "STUB: not implemented"
	return false
}

// reconnectDeadEndConnections disonnects connections to processes which are
// not in the set of processes to be run, and, if an out-port for a process
// supposed to be run gets disconnected, its out-port(s) will be connected to
// the sink instead, to make sure it is properly executed.
func (wf *Workflow) reconnectDeadEndConnections(procs map[string]WorkflowProcess) {
	_ = "STUB: not implemented"
	return
}

// OutPorts

// If the remotely connected process is not among the ones to run ...

// OutParamPorts

// If the remotely connected process is not among the ones to run ...

// Allow for a workflow with a single process
// A process can't both be the driver and be included in the main procs
// map, so if we have an alerative driver, it should not be in the main
// procs map

// upstreamProcsForProc returns all processes it is connected to, either
// directly or indirectly, via its in-ports and param-in-ports
func upstreamProcsForProc(proc WorkflowProcess) map[string]WorkflowProcess {
	_ = "STUB: not implemented"
	return nil
}

func mergeWFMaps(a map[string]WorkflowProcess, b map[string]WorkflowProcess) map[string]WorkflowProcess {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow) Auditf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (wf *Workflow) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (wf *Workflow) Fail(msg interface{}) { _ = "STUB: not implemented"; return }
