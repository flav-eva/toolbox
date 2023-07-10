package xlog

// ILogger is a generic logging interface
type ILogger interface {
	// Init options
	Init(options ...OptionFunc) error
	// Options for Logger
	Options() Options
	// Fields for Logger
	Fields(fields map[string]any) ILogger
	// Name returns the name of the xlog
	Name() string
	// Log write a log entry
	Log(level Level, v ...any)
	Logf(level Level, format string, v ...any)
}

// Level *********************** RAFAELLLE ****************** //
type Level int8

const (
	TraceLevel Level = iota - 1
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

// DefaultLogger *********************** RAFAELLLE ****************** //
var DefaultLogger *Logger

func Debug(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(DebugLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(DebugLevel, args...)
}

func Debugf(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(DebugLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(DebugLevel, format, args...)
}

func Info(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(InfoLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(InfoLevel, args...)
}

func Infof(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(InfoLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(InfoLevel, format, args...)
}

func Warn(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(WarnLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(WarnLevel, args...)
}

func Warnf(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(WarnLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(WarnLevel, format, args...)
}

func Error(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(ErrorLevel, args...)
}

func Errorf(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(ErrorLevel, format, args...)
}

func Panic(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(PanicLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(PanicLevel, args...)
}

func Panicf(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(PanicLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(PanicLevel, format, args...)
}

func Fatal(args ...any) {
	if !DefaultLogger.Options().Level.Enabled(FatalLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Log(FatalLevel, args...)
}

func Fatalf(format string, args ...any) {
	if !DefaultLogger.Options().Level.Enabled(FatalLevel) {
		return
	}
	DefaultLogger.Fields(DefaultLogger.fields).Logf(FatalLevel, format, args...)
}
