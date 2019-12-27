package logger

// LogLevel represents the level at which to start logging from.
// For example, if Warning (or 2) is the log level, only Warning (2),
// Error (3) and Fatal (4) logs will be captured.
type LogLevel int

// Log levels
const (
	Debug   LogLevel = iota // 0
	Info                    // 1
	Warning                 // 2
	Error                   // 3
	Fatal                   // 4
)

// Log level prefixes
var levelPrefixes = []string{"[DEBU]", "[INFO]", "[WARN]", "[ERRO]", "[FATA]"}
