package scipipe

import (
	"bytes"
	"os"
	"sync"
)

// IP Is the base interface which all other IPs need to adhere to
type IP interface {
	ID() string
	FinalizePath()
}

// ------------------------------------------------------------------------
// BaseIP type
// ------------------------------------------------------------------------

// BaseIP contains foundational functionality which all IPs need to implement.
// It is meant to be embedded into other IP implementations.
type BaseIP struct {
	path      string
	id        string
	auditInfo *AuditInfo
}

// NewBaseIP creates a new BaseIP
func NewBaseIP(path string) *BaseIP { _ = "STUB: not implemented"; return nil }

// ID returns a globally unique ID for the IP
func (ip *BaseIP) ID() string {
	_ = "STUB: not implemented"

	// ------------------------------------------------------------------------
	// FileIP type
	// ------------------------------------------------------------------------
	return ""
}

// FileIP (Short for "Information Packet" in Flow-Based Programming terminology)
// contains information and helper methods for a physical file on a normal disk.
type FileIP struct {
	*BaseIP
	buffer    *bytes.Buffer
	doStream  bool
	lock      *sync.Mutex
	SubStream *InPort
}

// NewFileIP creates a new FileIP
func NewFileIP(path string) (*FileIP, error) { _ = "STUB: not implemented"; return nil, nil }

// This will populate the audit info from file

//Don't init buffer if not needed?
//buf := make([]byte, 0, 128)
//ip.buffer = bytes.NewBuffer(buf)

func pathIsValid(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ------------------------------------------------------------------------
// Path stuff
// ------------------------------------------------------------------------

// Path returns the (final) path of the physical file
func (ip *FileIP) Path() string {
	_ = "STUB: not implemented"

	// TempDir returns the path to a temporary directory where outputs are written
	return ""
}

func (ip *FileIP) TempDir() string { _ = "STUB: not implemented"; return "" }

// TempPath returns the temporary path of the physical file
func (ip *FileIP) TempPath() string { _ = "STUB: not implemented"; return "" }

// FSRootPlaceHolder is a string to use instead of an initial '/', to indicate
// a path that belongs to the absolute root
const FSRootPlaceHolder = "__fsroot__"

// FifoPath returns the path to use when a FIFO file is used instead of a
// normal file
func (ip *FileIP) FifoPath() string { _ = "STUB: not implemented"; return "" }

// ------------------------------------------------------------------------
// Check-thing stuff
// ------------------------------------------------------------------------

// Size returns the size of an existing file, in bytes
func (ip *FileIP) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Exists checks if the file exists (at its final file name)
func (ip *FileIP) Exists() bool { _ = "STUB: not implemented"; return false }

// TempFileExists checks if the temp-file exists
func (ip *FileIP) TempFileExists() bool { _ = "STUB: not implemented"; return false }

// FifoFileExists checks if the FIFO-file (named pipe file) exists
func (ip *FileIP) FifoFileExists() bool { _ = "STUB: not implemented"; return false }

func (ip *FileIP) String() string {
	_ = "STUB: not implemented"

	// ------------------------------------------------------------------------
	// Open file-stuff
	// ------------------------------------------------------------------------
	return ""
}

// Open opens the file and returns a file handle (*os.File)
func (ip *FileIP) Open() *os.File { _ = "STUB: not implemented"; return nil }

// OpenTemp opens the temp file and returns a file handle (*os.File)
func (ip *FileIP) OpenTemp() *os.File { _ = "STUB: not implemented"; return nil }

// ------------------------------------------------------------------------
// FIFO-specific stuff
// ------------------------------------------------------------------------

// CreateFifo creates a FIFO file for the FileIP
func (ip *FileIP) CreateFifo() { _ = "STUB: not implemented"; return }

// RemoveFifo removes the FIFO file, if it exists
func (ip *FileIP) RemoveFifo() {
	_ = "STUB: not implemented"
	// FIXME: Shouldn't we check first whether the fifo exists?
	return
}

// ------------------------------------------------------------------------
// Read/Write stuff
// ------------------------------------------------------------------------

// Read reads the whole content of the file and returns the content as a byte
// array
func (ip *FileIP) Read() []byte { _ = "STUB: not implemented"; return nil }

// Write writes a byte array ([]byte) to the file's temp file path
func (ip *FileIP) Write(dat []byte) { _ = "STUB: not implemented"; return }

const (
	finalizePathMaxTries      = 3
	finalizePathBackoffFactor = 4
)

// FinalizePath renames the temporary file name to the final file name, thus enabling
// to separate unfinished, and finished files
func (ip *FileIP) FinalizePath() { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// Params and tags
// ------------------------------------------------------------------------

// Param returns the parameter named key, from the IPs audit info
func (ip *FileIP) Param(key string) string { _ = "STUB: not implemented"; return "" }

// ------------------------------------------------------------------------
// Tags stuff
// ------------------------------------------------------------------------

// Tag returns the tag for the tag with key k from the IPs audit info
func (ip *FileIP) Tag(k string) string { _ = "STUB: not implemented"; return "" }

// Tags returns the audit info's tags
func (ip *FileIP) Tags() map[string]string { _ = "STUB: not implemented"; return nil }

// AddTag adds the tag k with value v
func (ip *FileIP) AddTag(k string, v string) { _ = "STUB: not implemented"; return }

// AddTags adds a map of tags to the IPs audit info
func (ip *FileIP) AddTags(tags map[string]string) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// AuditInfo stuff
// ------------------------------------------------------------------------

// AuditFilePath returns the file path of the audit info file for the FileIP
func (ip *FileIP) AuditFilePath() string { _ = "STUB: not implemented"; return "" }

// SetAuditInfo sets the AuditInfo struct for the FileIP
func (ip *FileIP) SetAuditInfo(ai *AuditInfo) { _ = "STUB: not implemented"; return }

// WriteAuditLogToFile writes the audit log to its designated file
func (ip *FileIP) WriteAuditLogToFile() { _ = "STUB: not implemented"; return }

// AuditInfo returns the AuditInfo struct for the FileIP
func (ip *FileIP) AuditInfo() *AuditInfo { _ = "STUB: not implemented"; return nil }

// UnmarshalAuditInfoJSONFile returns an AuditInfo object from an AuditInfo
// .json file
func UnmarshalAuditInfoJSONFile(fileName string) (auditInfo *AuditInfo) {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------------------
// Extra convenience functions
// ------------------------------------------------------------------------

// UnMarshalJSON is a helper function to unmarshal the content of the IPs file
// to the interface v
func (ip *FileIP) UnMarshalJSON(v interface{}) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------
// Helper functions
// ------------------------------------------------------------------------

func (ip *FileIP) Failf(msg string, parts ...interface{}) { _ = "STUB: not implemented"; return }

func (ip *FileIP) Fail(msg interface{}) { _ = "STUB: not implemented"; return }

// CreateDirs creates all directories needed to enable writing the IP to its
// path (or temporary-path). If baseDir is provided, it will be prepended
// before all the IPs own temp path. This is to allow components to create their
// own temporary directory, to create the tasks in.
func (ip *FileIP) createDirs(baseDir string) { _ = "STUB: not implemented"; return }

func sanitizePathFragment(s string) (sanitized string) { _ = "STUB: not implemented"; return "" }
