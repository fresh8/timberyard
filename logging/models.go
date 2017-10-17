package logging

import "github.com/sirupsen/logrus"

// Opts configures various logging options
type Opts struct {
	ServiceName  string
	ServiceGroup string
	Level        string

	Hooks []logrus.Hook

	Formatter logrus.Formatter
}
