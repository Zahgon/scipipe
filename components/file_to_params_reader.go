package components

import (
	"github.com/scipipe/scipipe"
)

// FileToParamsReader takes a file path on its FilePath in-port, and returns the file
// content as []byte on its out-port Out
type FileToParamsReader struct {
	scipipe.BaseProcess
	filePath string
}

// NewFileToParamsReader returns an initialized new FileToParamsReader
func NewFileToParamsReader(wf *scipipe.Workflow, name string, filePath string) *FileToParamsReader {
	_ = "STUB: not implemented"
	return nil
}

// OutLine returns an parameter out-port with lines of the files being read
func (p *FileToParamsReader) OutLine() *scipipe.OutParamPort { _ = "STUB: not implemented"; return nil }

// Run the FileToParamsReader
func (p *FileToParamsReader) Run() { _ = "STUB: not implemented"; return }
