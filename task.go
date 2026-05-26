package scipipe

import (
	"time"
)

// Task represents a single static shell command, or go function, to be
// executed, and are scheduled and managed by a corresponding Process
type Task struct {
	Name          string
	Command       string
	CustomExecute func(*Task)
	InIPs         map[string]*FileIP
	OutIPs        map[string]*FileIP
	Params        map[string]string
	Tags          map[string]string
	Done          chan int
	cores         int
	workflow      *Workflow
	Process       *Process
	portInfos     map[string]*PortInfo
	subStreamIPs  map[string][]*FileIP
}

// ------------------------------------------------------------------------
// Factory method(s)
// ------------------------------------------------------------------------

// NewTask instantiates and initializes a new Task
func NewTask(workflow *Workflow, process *Process, name string, cmdPat string, inIPs map[string]*FileIP, outPathFuncs map[string]func(*Task) string, portInfos map[string]*PortInfo, params map[string]string, tags map[string]string, prepend string, customExecute func(*Task), cores int) *Task {
	_ = "STUB: not implemented"
	return nil
}

// Collect substream IPs

// Merge multiple input paths from a substream on the IP, into one string

// Create Out-IPs

const (
	parentDirPlaceHolder = "__parent__"
)

// formatCommand is a helper function for NewTask, that formats a shell command
// based on concrete file paths and parameter values
func (t *Task) formatCommand(cmd string, portInfos map[string]*PortInfo, inIPs map[string]*FileIP, subStreamIPs map[string][]*FileIP, outIPs map[string]*FileIP, params map[string]string, tags map[string]string, prepend string) string {
	_ = "STUB: not implemented"
	return ""
}

// Merge multiple input paths from a substream on the IP, into one string

// Add prepend string to the command

// ------------------------------------------------------------------------
// Main API methods: Accessor methods
// ------------------------------------------------------------------------

// InIP returns an IP for the in-port with name portName
func (t *Task) InIP(portName string) *FileIP { _ = "STUB: not implemented"; return nil }

// InPath returns the path name of an input file for the task
func (t *Task) InPath(portName string) string { _ = "STUB: not implemented"; return "" }

// OutIP returns an IP for the in-port with name portName
func (t *Task) OutIP(portName string) *FileIP { _ = "STUB: not implemented"; return nil }

// OutPath returns the path name of an input file for the task
func (t *Task) OutPath(portName string) string { _ = "STUB: not implemented"; return "" }

// Param returns the value of a param, for the task
func (t *Task) Param(portName string) string { _ = "STUB: not implemented"; return "" }

// Tag returns the value of a param, for the task
func (t *Task) Tag(tagName string) string { _ = "STUB: not implemented"; return "" }

// ------------------------------------------------------------------------
// Execute the task
// ------------------------------------------------------------------------

// Execute executes the task (the shell command or go function in CustomExecute)
func (t *Task) Execute() {
	_ = "STUB: not implemented"

	// Do some sanity checks
	return
}

// Execute task
// Will block if max concurrent tasks is reached
// Create output directories needed for any outputs

// ------------------------------------------------------------------------
// Helper methods for the Execute method
// ------------------------------------------------------------------------

// anyTempFileExists checks if any temporary workflow files exist and if so, returns true
func (t *Task) tempDirsExist() bool { _ = "STUB: not implemented"; return false }

// anyOutputsExist if any output file IP, or temporary file IPs, exist
func (t *Task) anyOutputsExist() (anyFileExists bool) { _ = "STUB: not implemented"; return false }

// createDirs creates directories for out-IPs of the task
func (t *Task) createDirs() error { _ = "STUB: not implemented"; return nil }

// This will create all out dirs, including the temp dir
// Temp dirs are not created for fifo files

// executeCommand executes the shell command cmd via bash
func (t *Task) executeCommand(cmd string) {
	_ = "STUB: not implemented"
	// cd into the task's tempdir, execute the command, and cd back
	return
}

func (t *Task) writeAuditLogs(startTime time.Time, finishTime time.Time) {
	_ = "STUB: not implemented"
	// Append audit info for the task to all its output IPs
	return
}

// Set the audit infos from incoming IPs into the "Upstream" map

// Add output paths generated for this task

// Add the current audit info to output ips and write them to file

func (t *Task) ensureAllOutputsExist() { _ = "STUB: not implemented"; return }

func (t *Task) finalizePaths() error { _ = "STUB: not implemented"; return nil }

func (t *Task) Auditf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (t *Task) Audit(msg string) { _ = "STUB: not implemented"; return }

func (t *Task) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (t *Task) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

// FinalizePaths renames temporary output files/directories to their proper paths.
// It is called both from Task, and from Process that implement cutom execution
// schedule.
func FinalizePaths(tempExecDir string, ips ...*FileIP) error { _ = "STUB: not implemented"; return nil }

// Move paths for ports, to final destinations

// For remaining paths in temporary execution dir, just move out of it

// Remove temporary execution dir (but not for absolute paths, or current dir)

var tempDirPrefix = "_scipipe_tmp"

// TempDir returns a string that is unique to a task, suitable for use
// in file paths. It is built up by merging all input filenames and parameter
// values that a task takes as input, joined with dots.
func (t *Task) TempDir() string { _ = "STUB: not implemented"; return "" }

// If resulting name is longer than 255

func prependParentDirPath(path string) string { _ = "STUB: not implemented"; return "" }

// For relative paths, add ".." to get out of current dir

func replaceParentDirsWithPlaceholder(pathSegment string) string {
	_ = "STUB: not implemented"
	return ""
}

func replacePlaceholdersWithParentDirs(pathSegment string) string {
	_ = "STUB: not implemented"
	return ""
}
