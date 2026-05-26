package scipipe

import (
	"time"
)

// AuditInfo contains structured audit/provenance logging information for a
// particular task (invocation), to go with all outgoing IPs from that task
type AuditInfo struct {
	ID          string
	ProcessName string
	Command     string
	Params      map[string]string
	Tags        map[string]string
	StartTime   time.Time
	FinishTime  time.Time
	ExecTimeNS  time.Duration
	OutFiles    map[string]string
	Upstream    map[string]*AuditInfo
}

// NewAuditInfo returns a new AuditInfo struct
func NewAuditInfo() *AuditInfo { _ = "STUB: not implemented"; return nil }
