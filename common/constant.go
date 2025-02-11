package common

type LogLevel string

type SortDir string

const (
	LogLevelError   LogLevel = "error"
	LogLevelWarning LogLevel = "warn"
	LogLevelTrace   LogLevel = "trace"
	LogLevelInfo    LogLevel = "info"

	SortDirAsc  SortDir = "ASC"
	SortDirDesc SortDir = "DESC"

	AuditLogActionUpdate = "UPDATE"
	AuditLogActionDelete = "DELETE"
)
