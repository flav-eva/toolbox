package xlog

import (
	"context"
	"io"
)

// Options 可选项，可选字段，xlog 一些基础配置
type Options struct {
	// xlog name
	Name string
	// logging level. default is InfoLevel
	Level Level
	// fields to log
	Fields map[string]any
	// output file, default is os.Stdout
	Out io.Writer
	// Caller skip frame file:line info
	CallerSkipCount int
	// context
	Context context.Context
}

type OptionFunc func(*Options)

// WithLevel 等同于 SetLevel for Options
//
//	func(o *Options, level Level) {
//	  o.Level = level
//	}
func WithLevel(level Level) OptionFunc {
	return func(o *Options) {
		o.Level = level
	}
}

// WithFields set fields for Options
func WithFields(fields map[string]any) OptionFunc {
	return func(o *Options) {
		o.Fields = fields
	}
}

func WithOutput(out io.Writer) OptionFunc {
	return func(o *Options) {
		o.Out = out
	}
}

func WithCallerSkipCount(skip int) OptionFunc {
	return func(o *Options) {
		o.CallerSkipCount = skip
	}
}

func WithName(name string) OptionFunc {
	return func(o *Options) {
		o.Name = name
	}
}

func SetCtxValue(k, v any) OptionFunc {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
