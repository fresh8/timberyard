package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uber-common/bark"
)

// Fields is an alias for bark.Fields
type Fields = bark.Fields

// Log is the importable logger
var (
	logger bark.Logger
	serviceName string
	serviceGroup string
	host string
	LogrusLogger *logrus.Logger
)

// Initialise sets the logger
func Initialise(opts Opts) {
	if IsInitialised() {
		Log().Warn("logger has already been initialised")
		return
	}

	serviceName = opts.ServiceName
	serviceGroup = opts.ServiceGroup
	host, _ = os.Hostname()

	LogrusLogger = logrus.New()
	if opts.Formatter == nil {
		opts.Formatter = &logrus.JSONFormatter{}
	}
	LogrusLogger.Formatter = opts.Formatter

	level, err := logrus.ParseLevel(opts.Level)
	if err == nil {
		LogrusLogger.Level = level
	} else {
		LogrusLogger.Level = logrus.ErrorLevel
	}

	for _, hook := range opts.Hooks {
		LogrusLogger.Hooks.Add(hook)
	}

	logger = bark.NewLoggerFromLogrus(LogrusLogger)
}

// Use initialises the logger package from an existing generic logger
func Use(newLogger bark.Logger) {
	logger = newLogger
}

// Logger returns the logger object for use
func Logger() bark.Logger {
	return logger
}

// IsInitialised returns if the logger is set up correctly
func IsInitialised() bool {
	return logger != nil
}

// Log wraps logrus logging with the service fields
func Log() bark.Logger {
	if !IsInitialised() {
		logrus.Fatal("The logging package must be initialised")
	}

	return logger.WithFields(bark.Fields{
		"service": serviceName,
		"group":   serviceGroup,
		"host":    host,
		"pid":     os.Getpid(),
	})
}

// WithFields is an alias for logrus.WithFields to accept our own
func WithFields(f Fields) bark.Logger {
	lf := bark.Fields(f)
	return Log().WithFields(lf)
}
