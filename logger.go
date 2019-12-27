package logger

import (
	"fmt"
	"io"
	"log"
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

func (l Logger) println(level LogLevel, v ...interface{}) {
	l.log.SetPrefix(levelPrefixes[l.level])
	l.log.Println(fmt.Sprintf("%s %s", levelPrefixes[level], fmt.Sprint(v...)))
}

// Debug ...
func (l Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.println(Debug, v...)
}

// // Info ...
// func (l Logger) Info() {
// 	if Info < l.level {
// 		return
// 	}
// }

// // Warning ...
// func (l Logger) Warning() {
// 	if Warning < l.level {
// 		return
// 	}
// }

// // Error ...
// func (l Logger) Error() {
// 	if Error < l.level {
// 		return
// 	}
// }

// // Fatal ...
// func (l Logger) Fatal() {
// 	if Fatal < l.level {
// 		return
// 	}
// }
