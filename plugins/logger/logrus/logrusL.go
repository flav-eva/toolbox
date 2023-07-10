package logrus

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/flav-eva/toolbox/xlog"
)

type entryLogger interface {
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Log(level logrus.Level, args ...interface{})
	Logf(level logrus.Level, format string, args ...interface{})
}

type logrusLogger struct {
	Logger entryLogger
	opts   Options
}

func NewLogrusLogger(opts ...xlog.OptionFunc) xlog.ILogger {
	options := Options{
		Options: xlog.Options{
			Level:   xlog.InfoLevel,
			Fields:  make(map[string]interface{}),
			Out:     os.Stderr,
			Context: context.Background(),
		},
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		ReportCaller: false,
		ExitFunc:     os.Exit,
	}
	l := &logrusLogger{opts: options}
	_ = l.Init(opts...)
	return l
}

func (l *logrusLogger) Init(opts ...xlog.OptionFunc) error {
	for _, o := range opts {
		o(&l.opts.Options)
	}

	if formatter, ok := l.opts.Context.Value(formatterKey{}).(logrus.Formatter); ok {
		l.opts.Formatter = formatter
	}
	if hs, ok := l.opts.Context.Value(hooksKey{}).(logrus.LevelHooks); ok {
		l.opts.Hooks = hs
	}
	if caller, ok := l.opts.Context.Value(reportCallerKey{}).(bool); ok && caller {
		l.opts.ReportCaller = caller
	}
	if exitFunction, ok := l.opts.Context.Value(exitKey{}).(func(int)); ok {
		l.opts.ExitFunc = exitFunction
	}

	switch ll := l.opts.Context.Value(logrusLoggerKey{}).(type) {
	case *logrus.Logger:
		// overwrite default options
		l.opts.Level = logrusToLoggerLevel(ll.GetLevel())
		l.opts.Out = ll.Out
		l.opts.Formatter = ll.Formatter
		l.opts.Hooks = ll.Hooks
		l.opts.ReportCaller = ll.ReportCaller
		l.opts.ExitFunc = ll.ExitFunc
		l.Logger = ll
	case *logrus.Entry:
		// overwrite default options
		el := ll.Logger
		l.opts.Level = logrusToLoggerLevel(el.GetLevel())
		l.opts.Out = el.Out
		l.opts.Formatter = el.Formatter
		l.opts.Hooks = el.Hooks
		l.opts.ReportCaller = el.ReportCaller
		l.opts.ExitFunc = el.ExitFunc
		l.Logger = ll
	case nil:
		log := logrus.New() // defaults
		log.SetLevel(loggerToLogrusLevel(l.opts.Level))
		log.SetOutput(l.opts.Out)
		log.SetFormatter(l.opts.Formatter)
		log.ReplaceHooks(l.opts.Hooks)
		log.SetReportCaller(l.opts.ReportCaller)
		log.ExitFunc = l.opts.ExitFunc
		l.Logger = log
	default:
		return fmt.Errorf("invalid logrus type: %T", ll)
	}

	return nil
}

func (l *logrusLogger) Name() string {
	return "logrus"
}

func (l *logrusLogger) Fields(fields map[string]interface{}) xlog.ILogger {
	return &logrusLogger{l.Logger.WithFields(fields), l.opts}
}

func (l *logrusLogger) Log(level xlog.Level, args ...interface{}) {
	l.Logger.Log(loggerToLogrusLevel(level), args...)
}

func (l *logrusLogger) Logf(level xlog.Level, format string, args ...interface{}) {
	l.Logger.Logf(loggerToLogrusLevel(level), format, args...)
}

func (l *logrusLogger) Options() xlog.Options {
	return l.opts.Options
}

func loggerToLogrusLevel(level xlog.Level) logrus.Level {
	switch level {
	case xlog.TraceLevel:
		return logrus.TraceLevel
	case xlog.DebugLevel:
		return logrus.DebugLevel
	case xlog.InfoLevel:
		return logrus.InfoLevel
	case xlog.WarnLevel:
		return logrus.WarnLevel
	case xlog.ErrorLevel:
		return logrus.ErrorLevel
	case xlog.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func logrusToLoggerLevel(level logrus.Level) xlog.Level {
	switch level {
	case logrus.TraceLevel:
		return xlog.TraceLevel
	case logrus.DebugLevel:
		return xlog.DebugLevel
	case logrus.InfoLevel:
		return xlog.InfoLevel
	case logrus.WarnLevel:
		return xlog.WarnLevel
	case logrus.ErrorLevel:
		return xlog.ErrorLevel
	case logrus.FatalLevel:
		return xlog.FatalLevel
	default:
		return xlog.InfoLevel
	}
}
