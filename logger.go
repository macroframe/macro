package macro

import (
	"github.com/pterm/pterm"
)

type loggerType struct {
	logger pterm.Logger
}

// Logger includes all the methods for logging.
var Logger = loggerType{
	logger: pterm.DefaultLogger,
}

// Info logs an info message.
func (l loggerType) Info(message string, args ...any) {
	l.logger.Info(message, l.logger.Args(args...))
}

// Warn logs a warning message.
func (l loggerType) Warn(message string, args ...any) {
	l.logger.Warn(message, l.logger.Args(args...))
}

// Error logs an error message.
func (l loggerType) Error(message string, args ...any) {
	l.logger.Error(message, l.logger.Args(args...))
}

// Fatal logs a fatal message and exits the program.
func (l loggerType) Fatal(message string, args ...any) {
	l.logger.Fatal(message, l.logger.Args(args...))
}
