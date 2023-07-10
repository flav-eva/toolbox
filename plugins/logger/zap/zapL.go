package zap

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/flav-eva/toolbox/xlog"
)

type zapLog struct {
	cfg  zap.Config
	zap  *zap.Logger
	opts xlog.Options

	sync.RWMutex
	fields map[string]any
}

func NewZapLogger(opts ...xlog.OptionFunc) xlog.ILogger {
	options := xlog.Options{
		Level:           xlog.InfoLevel,
		Fields:          make(map[string]any),
		Out:             os.Stdout,
		CallerSkipCount: 3,
		Context:         context.Background(),
	}
	l := &zapLog{opts: options}
	if err := l.Init(opts...); err != nil {
		panic(err)
	}
	return l
}

func (z *zapLog) Init(opts ...xlog.OptionFunc) error {

	for _, opt := range opts {
		opt(&z.opts)
	}

	zapConfig := zap.NewProductionConfig()
	if zconfig, ok := z.opts.Context.Value(configKey{}).(zap.Config); ok {
		zapConfig = zconfig
	}
	if zeconfig, ok := z.opts.Context.Value(encoderConfigKey{}).(zapcore.EncoderConfig); ok {
		zapConfig.EncoderConfig = zeconfig
	}
	writer, ok := z.opts.Context.Value(writerKey{}).(io.Writer)
	if !ok {
		writer = os.Stdout
	}
	skip, ok := z.opts.Context.Value(callerSkipKey{}).(int)
	if !ok || skip < 1 {
		skip = 1
	}

	zapConfig.Level = zap.NewAtomicLevel()
	if z.opts.Level != xlog.InfoLevel {
		zapConfig.Level.SetLevel(loggerToZapLevel(z.opts.Level))
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer)),
		zapConfig.Level)

	log := zap.New(logCore, zap.AddCaller(), zap.AddCallerSkip(skip), zap.AddStacktrace(zap.DPanicLevel))
	if z.opts.Fields != nil {
		zFields := make([]zap.Field, 0)
		for k, v := range z.opts.Fields {
			zFields = append(zFields, zap.Any(k, v))
		}
		log = log.With(zFields...)
	}

	namespace, ok := z.opts.Context.Value(namespaceKey{}).(string)
	if ok {
		log = log.With(zap.Namespace(namespace))
	}

	z.cfg = zapConfig
	z.zap = log
	z.fields = make(map[string]any)

	return nil
}

func (z *zapLog) Fields(fields map[string]any) xlog.ILogger {
	z.Lock()
	nFields := make(map[string]any, len(z.fields))
	for k, v := range z.fields {
		nFields[k] = v
	}
	z.Unlock()
	for k, v := range fields {
		nFields[k] = v
	}
	return &zapLog{
		cfg:    z.cfg,
		zap:    z.zap,
		opts:   z.opts,
		fields: nFields,
	}
}

func (z *zapLog) Log(level xlog.Level, args ...any) {
	z.RLock()
	zFields := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		zFields = append(zFields, zap.Any(k, v))
	}
	z.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	switch lvl {
	case zap.DebugLevel:
		z.zap.Debug(msg, zFields...)
	case zap.InfoLevel:
		z.zap.Info(msg, zFields...)
	case zap.WarnLevel:
		z.zap.Warn(msg, zFields...)
	case zap.ErrorLevel:
		z.zap.Error(msg, zFields...)
	case zap.PanicLevel:
		z.zap.Panic(msg, zFields...)
	case zap.FatalLevel:
		z.zap.Fatal(msg, zFields...)
	}
}

func (z *zapLog) Logf(level xlog.Level, format string, args ...any) {
	z.Lock()
	zFields := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		zFields = append(zFields, zap.Any(k, v))
	}
	z.Unlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	switch lvl {
	case zap.DebugLevel:
		z.zap.Debug(msg, zFields...)
	case zap.InfoLevel:
		z.zap.Info(msg, zFields...)
	case zap.WarnLevel:
		z.zap.Warn(msg, zFields...)
	case zap.ErrorLevel:
		z.zap.Error(msg, zFields...)
	case zap.PanicLevel:
		z.zap.Panic(msg, zFields...)
	case zap.FatalLevel:
		z.zap.Fatal(msg, zFields...)
	}
}

func (z *zapLog) Name() string {
	return "zap"
}

func (z *zapLog) Options() xlog.Options {
	return z.opts
}

func loggerToZapLevel(level xlog.Level) zapcore.Level {
	switch level {
	case xlog.TraceLevel, xlog.DebugLevel:
		return zapcore.DebugLevel
	case xlog.InfoLevel:
		return zapcore.InfoLevel
	case xlog.WarnLevel:
		return zapcore.WarnLevel
	case xlog.ErrorLevel:
		return zapcore.ErrorLevel
	case xlog.PanicLevel:
		return zapcore.PanicLevel
	case xlog.FatalLevel:
		return zapcore.FatalLevel

	default:
		return zapcore.InfoLevel
	}
}

func zapToLoggerLevel(level zapcore.Level) xlog.Level {
	switch level {
	case zap.DebugLevel:
		return xlog.DebugLevel
	case zap.InfoLevel:
		return xlog.InfoLevel
	case zap.WarnLevel:
		return xlog.WarnLevel
	case zap.ErrorLevel:
		return xlog.ErrorLevel
	case zap.PanicLevel:
		return xlog.PanicLevel
	case zap.FatalLevel:
		return xlog.FatalLevel

	default:
		return xlog.InfoLevel
	}
}
