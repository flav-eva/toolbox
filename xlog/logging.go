package xlog

type Logger struct {
	ILogger
	fields map[string]any
}

// NewLogger 这里对外封装了一套依赖 ILogger 的接口，默认为 DefaultLogger
// 我们可以在项目里手动配置，根据plugin：logrus or zap
func NewLogger(log ILogger) *Logger {
	return &Logger{
		ILogger: log,
		fields:  make(map[string]any),
	}
}

func (l *Logger) WithFields(fields map[string]any) *Logger {
	nFields := copyFields(fields)
	for k, v := range l.fields {
		nFields[k] = v
	}
	return &Logger{ILogger: l, fields: nFields}
}

func (l *Logger) WithError(err error) *Logger {
	nFields := copyFields(l.fields)
	nFields["error"] = err.Error()
	return &Logger{ILogger: l, fields: nFields}
}

func (l *Logger) Trace(args ...any) {
	if !l.Options().Level.Enabled(TraceLevel) {
		return
	}
	l.Fields(l.fields).Log(TraceLevel, args...)
}

func (l *Logger) Tracef(format string, args ...any) {
	if !l.Options().Level.Enabled(TraceLevel) {
		return
	}
	l.Fields(l.fields).Logf(TraceLevel, format, args...)
}

func (l *Logger) Debug(args ...any) {
	if !l.Options().Level.Enabled(DebugLevel) {
		return
	}
	l.Fields(l.fields).Log(DebugLevel, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	if !l.Options().Level.Enabled(DebugLevel) {
		return
	}
	l.Fields(l.fields).Logf(DebugLevel, format, args...)
}

func (l *Logger) Info(args ...any) {
	if !l.Options().Level.Enabled(InfoLevel) {
		return
	}
	l.Fields(l.fields).Log(InfoLevel, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if !l.Options().Level.Enabled(InfoLevel) {
		return
	}
	l.Fields(l.fields).Logf(InfoLevel, format, args...)
}

func (l *Logger) Warn(args ...any) {
	if !l.Options().Level.Enabled(WarnLevel) {
		return
	}
	l.Fields(l.fields).Log(WarnLevel, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if !l.Options().Level.Enabled(WarnLevel) {
		return
	}
	l.Fields(l.fields).Logf(WarnLevel, format, args...)
}

func (l *Logger) Error(args ...any) {
	if !l.Options().Level.Enabled(ErrorLevel) {
		return
	}
	l.Fields(l.fields).Log(ErrorLevel, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if !l.Options().Level.Enabled(ErrorLevel) {
		return
	}
	l.Fields(l.fields).Logf(ErrorLevel, format, args...)
}

func (l *Logger) Panic(args ...any) {
	if !l.Options().Level.Enabled(PanicLevel) {
		return
	}
	l.Fields(l.fields).Log(PanicLevel, args...)
}

func (l *Logger) Panicf(format string, args ...any) {
	if !l.Options().Level.Enabled(PanicLevel) {
		return
	}
	l.Fields(l.fields).Logf(PanicLevel, format, args...)
}

func (l *Logger) Fatal(args ...any) {
	if !l.Options().Level.Enabled(FatalLevel) {
		return
	}
	l.Fields(l.fields).Log(FatalLevel, args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if !l.Options().Level.Enabled(FatalLevel) {
		return
	}
	l.Fields(l.fields).Logf(FatalLevel, format, args...)
}
