package components

import (
	"os"

	"github.com/scipipe/scipipe"
)

// FileSplitter is a process that will split a file into multiple files, each
// with LinesPerSplit number of lines per file
type FileSplitter struct {
	scipipe.BaseProcess
	LinesPerSplit int
}

// NewFileSplitter returns an initialized FileSplitter process that will split a
// file into multiple files, each with linesPerSplit number of lines per file
func NewFileSplitter(wf *scipipe.Workflow, name string, linesPerSplit int) *FileSplitter {
	_ = "STUB: not implemented"
	return nil
}

// InFile returns the port for the input file
func (p *FileSplitter) InFile() *scipipe.InPort { _ = "STUB: not implemented"; return nil }

// OutSplitFile returns the resulting split (part) files generated0
func (p *FileSplitter) OutSplitFile() *scipipe.OutPort { _ = "STUB: not implemented"; return nil }

// Run runs the FileSplitter process
func (p *FileSplitter) Run() { _ = "STUB: not implemented"; return }

// Create new IP

func (p *FileSplitter) createNewSplitFile(ip *scipipe.FileIP, basePath string) (tempDir string, tempFile *os.File) {
	_ = "STUB: not implemented"
	return "", nil
}

var chars = []rune("abcdefghijklmnopqrstuvwxyz")

func getRandString(n int) string { _ = "STUB: not implemented"; return "" }

func (p *FileSplitter) newSplitIPFromIndex(basePath string, splitIdx int) *scipipe.FileIP {
	_ = "STUB: not implemented"
	return nil
}
