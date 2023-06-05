package macro

import (
	"github.com/pterm/pterm"
)

type logger struct {
	logger pterm.Logger
}

// Logger includes all the methods for logging.
var Logger = logger{
	logger: pterm.DefaultLogger,
}

// Info logs an info message.
func (l logger) Info(message string, args ...any) {
	l.logger.Info(message, l.logger.Args(args...))
}

// Warn logs a warning message.
func (l logger) Warn(message string, args ...any) {
	l.logger.Warn(message, l.logger.Args(args...))
}

// Error logs an error message.
func (l logger) Error(message string, args ...any) {
	l.logger.Error(message, l.logger.Args(args...))
}

// Fatal logs a fatal message and exits the program.
func (l logger) Fatal(message string, args ...any) {
	l.logger.Fatal(message, l.logger.Args(args...))
}
