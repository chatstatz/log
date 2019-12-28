package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger ...
type Logger struct {
	level LogLevel
	log   *log.Logger
}

// New returns a new Logger. It panics if an invalid
// level is provided. Valid levels include: debug, info,
// warning, error and fatal.
func New(level string, out io.Writer) *Logger {
	logLevel := GetLogLevelFromString(level)

	return &Logger{
		level: logLevel,
		log:   log.New(out, "", log.LstdFlags),
	}
}

func (l *Logger) println(level LogLevel, v ...interface{}) {
	l.log.Println(fmt.Sprintf("%s %s", levelPrefixes[level], fmt.Sprint(v...)))
}

// Debug prints debug log messages.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.println(Debug, v...)
}

// Debugf is the same as Debug but allows for fomatted debug messages.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.println(Debug, fmt.Sprintf(format, v...))
}

// Info prints informational log messages.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}

	l.println(Info, v...)
}

// Infof is the same as Info but allows for fomatted informational messages.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.println(Info, fmt.Sprintf(format, v...))
}

// Warn prints warning log messages.
func (l *Logger) Warn(v ...interface{}) {
	if Warning < l.level {
		return
	}

	l.println(Warning, v...)
}

// Warnf is the same as Warn but allows for fomatted warning messages.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warning < l.level {
		return
	}

	l.println(Warning, fmt.Sprintf(format, v...))
}

// Error prints error log messages.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.println(Error, v...)
}

// Errorf is the same as Error but allows for formatted error messsages.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.println(Error, fmt.Sprintf(format, v...))
}

// Fatal prints fatal error log messages. It calls
// os.Exit after logging.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.println(Fatal, v...)
	os.Exit(1)
}

// Fatalf is the same as Fatal but allows for fomatted errors.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.println(Fatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}
