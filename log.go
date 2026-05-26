package scipipe

import (
	"io"
	"log"
)

var (
	// Trace is a log handler for extremely detailed level logs. It is so far
	// sparely used in scipipe.
	Trace *log.Logger
	// Debug is a log handler for debugging level logs
	Debug *log.Logger
	// Info is a log handler for information level logs
	Info *log.Logger
	// Audit is a log handler for audit level logs
	Audit *log.Logger
	// Warning is a log handler for warning level logs
	Warning *log.Logger
	// Error is a log handler for error level logs
	Error     *log.Logger
	logExists bool
)

// InitLog initiates logging handlers
func InitLog(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	auditHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	_ = "STUB: not implemented"
	return
}

// This level is the one suggested to use when running scientific workflows, to retain audit
// information

// InitLogDebug initiates logging with level=DEBUG
func InitLogDebug() { _ = "STUB: not implemented"; return }

// InitLogInfo initiates logging with level=INFO
func InitLogInfo() { _ = "STUB: not implemented"; return }

// InitLogAudit initiate logging with level=AUDIT
func InitLogAudit() { _ = "STUB: not implemented"; return }

// InitLogAuditToFile initiate logging with level=AUDIT, and write that to
// fileName
func InitLogAuditToFile(filePath string) { _ = "STUB: not implemented"; return }

// InitLogWarning initiates logging with level=WARNING
func InitLogWarning() { _ = "STUB: not implemented"; return }

// InitLogError initiates logging with level=ERROR
func InitLogError() { _ = "STUB: not implemented"; return }
