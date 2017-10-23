package logging

import (
	"os"

	"github.com/fresh8/timberyard/mock_bark"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/uber-common/bark"
)

// Fields is an alias for logrus.Fields to keep dem imports tidy like
type Fields map[string]interface{}

// Log is the importable logger
var logger bark.Logger
var serviceName string
var serviceGroup string
var host string

// Initialise sets the logger
func Initialise(opts Opts) {
	if IsInitialised() {
		Log().Warn("logger has already been initialised")
		return
	}

	serviceName = opts.ServiceName
	serviceGroup = opts.ServiceGroup
	host, _ = os.Hostname()

	l := logrus.New()
	if opts.Formatter == nil {
		opts.Formatter = &logrus.JSONFormatter{}
	}
	l.Formatter = opts.Formatter

	level, err := logrus.ParseLevel(opts.Level)
	if err == nil {
		l.Level = level
	} else {
		l.Level = logrus.ErrorLevel
	}

	for _, hook := range opts.Hooks {
		l.Hooks.Add(hook)
	}

	logger = bark.NewLoggerFromLogrus(l)
}

func InitialiseMock(ctrl *gomock.Controller) {
	logger = mock_bark.NewMockLogger(ctrl)
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
