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

// New returns a new Logger. It panics if level is
// less than 0 or greater than 4.
func New(level LogLevel, out io.Writer) *Logger {
	if level < Debug || level > Fatal {
		panic(fmt.Sprintf("logger: %d is not a valid log level", level))
	}

	return &Logger{
		level: level,
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

// Error prints error log messages.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.println(Error, v...)
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
