package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// LogLevel represents the severity of a log message
type LogLevel int

// various log levels
const (
	FATAL LogLevel = iota
	ERROR
	WARN
	INFO
	TRACE
)

var currentLevel = INFO

// SetLogLevel changes the current log level
func SetLogLevel(level LogLevel) {
	currentLevel = level
}

// Log a function's output at a various level, ignoring those below the
// currently configured level.
func loggerf(level LogLevel, format string, args ...interface{}) {
	if level > currentLevel {
		return
	}

	var prefix string
	switch level {
	case FATAL:
		prefix = "FATAL"
	case ERROR:
		prefix = "ERROR"
	case WARN:
		prefix = "WARNING"
	case INFO:
		prefix = "INFO"
	case TRACE:
		prefix = "TRACE"
	}

	ptr, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Printf("- %s - %s\n", prefix, args)
	}

	fname := runtime.FuncForPC(ptr).Name()
	lastDot := 0
	for i := 0; i < len(fname); i++ {
		if fname[i] == '.' {
			lastDot = i
		}
	}
	if lastDot == 0 {
		fname = filepath.Base(fname)
	}

	preformatted := fmt.Sprintf(format, args...)
	log.Printf("- %s - %s:%d:%s() - %s",
		prefix, filepath.Base(file), line, fname[lastDot:], preformatted)
}

func logger(level LogLevel, args ...interface{}) {
	loggerf(level, "%s\n", args...)
}

// Fatalf logs and kills the program. Uses printf formatting.
func Fatalf(format string, args ...interface{}) {
	loggerf(FATAL, format, args...)
	os.Exit(1)
}

// Fatal logs and kills the program
func Fatal(args ...interface{}) {
	logger(FATAL, args...)
	os.Exit(1)
}

// Errorf logs at the Error level, but allows formatting
func Errorf(format string, args ...interface{}) {
	loggerf(ERROR, format, args...)
}

// Error logs at the Error level
func Error(args ...interface{}) {
	logger(ERROR, args...)
}

// Warnf logs at the Warn level, but allows formatting
func Warnf(format string, args ...interface{}) {
	loggerf(WARN, format, args...)
}

// Warn logs at the Warn level
func Warn(args ...interface{}) {
	logger(WARN, args)
}

// Infof logs at the Info level, but allows formatting
func Infof(format string, args ...interface{}) {
	loggerf(INFO, format, args...)
}

// Info logs at the Info level
func Info(args ...interface{}) {
	logger(INFO, args...)
}

// Tracef logs at the Warn level, but allows formatting
func Tracef(format string, args ...interface{}) {
	loggerf(TRACE, format, args...)
}

// Trace logs at the Trace level
func Trace(args ...interface{}) {
	logger(TRACE, args...)
}