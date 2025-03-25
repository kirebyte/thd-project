package logger

import (
	"fmt"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var currentLevel = INFO

// init is a special function that is called before the main function.
// It is used to set the log level based on the LOG_LEVEL environment variable.
// If the LOG_LEVEL is not set, the default level is INFO.
func init() {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "DEBUG":
		currentLevel = DEBUG
	case "INFO":
		currentLevel = INFO
	case "WARN":
		currentLevel = WARN
	case "ERROR":
		currentLevel = ERROR
	case "FATAL":
		currentLevel = FATAL
	}
}

// Debug logs a message with the DEBUG level.
func Debug(msg string) {
	if currentLevel <= DEBUG {
		logToStderr("[DEBUG] " + msg)
	}
}

// Info logs a message with the INFO level.
func Info(msg string) {
	if currentLevel <= INFO {
		logToStderr("[INFO] " + msg)
	}
}

// Warn logs a message with the WARN level.
func Warn(msg string) {
	if currentLevel <= WARN {
		logToStderr("[WARN] " + msg)
	}
}

// Error logs a message with the ERROR level.
func Error(msg string) {
	if currentLevel <= ERROR {
		logToStderr("[ERROR] " + msg)
	}
}

// Fatal logs a message with the FATAL level and exits the program with status code 1.
func Fatal(msg string) {
	if currentLevel <= FATAL {
		logToStderr("[FATAL] " + msg)
		os.Exit(1)
	}
}

// Fatalf logs a formatted message with the FATAL level and exits the program with status code 1.
func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

// logToStderr is a helper function that writes a message to the standard error output.
func logToStderr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}
