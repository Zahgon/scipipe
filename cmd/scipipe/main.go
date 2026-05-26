package main

import (
	"flag"
	"log"
	"os"
)

var (
	Info *log.Logger
)

func main() {
	initLogs()
	initHelp()
	flag.Parse()
	err := parseFlags(flag.Args())
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}

func parseFlags(args []string) error { _ = "STUB: not implemented"; return nil }

func parseArgsAudit2X(args []string, extension string) (inFile string, outFile string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func printNewUsage() { _ = "STUB: not implemented"; return }

func printAudit2HTMLUsage() { _ = "STUB: not implemented"; return }

func printHelp() { _ = "STUB: not implemented"; return }

func writeNewWorkflowFile(fileName string) { _ = "STUB: not implemented"; return }

func initHelp() { _ = "STUB: not implemented"; return }

func initLogs() { _ = "STUB: not implemented"; return }

func initLogsTest() { _ = "STUB: not implemented"; return }

const workflowStub = `// Workflow written in SciPipe.
// For more information about SciPipe, see: http://scipipe.org
package main

import sp "github.com/scipipe/scipipe"

func main() {
	// Create a workflow, using 4 cpu cores
	wf := sp.NewWorkflow("my_workflow", 4)

	// Initialize processes
	foo := wf.NewProc("fooer", "echo foo > {o:foo}")
	foo.SetOut("foo", "foo.txt")

	f2b := wf.NewProc("foo2bar", "sed 's/foo/bar/g' {i:foo} > {o:bar}")
	f2b.SetOut("bar", "{i:foo}.bar.txt")

	// From workflow dependency network
	f2b.In("foo").From(foo.Out("foo"))

	// Run the workflow
	wf.Run()
}`

func errWrap(err error, msg string) error { _ = "STUB: not implemented"; return nil }

func errWrapf(err error, msg string, v ...interface{}) error { _ = "STUB: not implemented"; return nil }
