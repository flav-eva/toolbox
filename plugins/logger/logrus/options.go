package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/flav-eva/toolbox/xlog"
)

type Options struct {
	xlog.Options
	Formatter logrus.Formatter
	Hooks     logrus.LevelHooks
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type formatterKey struct{}

func WithTextTextFormatter(formatter *logrus.TextFormatter) xlog.OptionFunc {
	return xlog.SetCtxValue(formatterKey{}, formatter)
}

func WithJSONFormatter(formatter *logrus.JSONFormatter) xlog.OptionFunc {
	return xlog.SetCtxValue(formatterKey{}, formatter)
}

type hooksKey struct{}

func WithLevelHooks(hooks logrus.LevelHooks) xlog.OptionFunc {
	return xlog.SetCtxValue(hooksKey{}, hooks)
}

type reportCallerKey struct{}

func ReportCaller() xlog.OptionFunc {
	return xlog.SetCtxValue(reportCallerKey{}, true)
}

type exitKey struct{}

func WithExitFunc(exit func(int)) xlog.OptionFunc {
	return xlog.SetCtxValue(exitKey{}, exit)
}

type logrusLoggerKey struct{}

func WithLogger(l logrus.StdLogger) xlog.OptionFunc {
	return xlog.SetCtxValue(logrusLoggerKey{}, l)
}
