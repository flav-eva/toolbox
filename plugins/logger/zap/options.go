package zap

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/flav-eva/toolbox/xlog"
)

type Options struct {
	xlog.Options
}

type callerSkipKey struct{}

func WithCallerSkip(i int) xlog.OptionFunc {
	return xlog.SetCtxValue(callerSkipKey{}, i)
}

type configKey struct{}

func WithConfig(c zap.Config) xlog.OptionFunc {
	return xlog.SetCtxValue(configKey{}, c)
}

type encoderConfigKey struct{}

func WithEncodeConfig(c zapcore.EncoderConfig) xlog.OptionFunc {
	return xlog.SetCtxValue(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) xlog.OptionFunc {
	return xlog.SetCtxValue(namespaceKey{}, namespace)
}

type writerKey struct{}

func WithOutput(out io.Writer) xlog.OptionFunc {
	return xlog.SetCtxValue(writerKey{}, out)
}
