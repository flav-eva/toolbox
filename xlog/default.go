package xlog

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func init() {
	// TODO 这里检查一些 conf 里的配置
	DefaultLogger = NewLogger(NewDefaultLogger())
}

type defaultLogger struct {
	sync.RWMutex
	opts Options
}

// Init (opts...) overwrite provided options
func (l *defaultLogger) Init(opts ...OptionFunc) error {
	l.Lock()
	defer l.Unlock()
	for _, opt := range opts {
		opt(&l.opts)
	}
	return nil
}

func (l *defaultLogger) Name() string {
	return l.opts.Name
}

func (l *defaultLogger) Options() Options {
	l.RLock()
	defer l.RUnlock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	return opts
}

func (l *defaultLogger) Fields(fields map[string]any) ILogger {
	l.Lock()
	defer l.Unlock()
	nFields := copyFields(fields)
	for k, v := range l.opts.Fields {
		nFields[k] = v
	}
	l.opts.Fields = nFields
	return l
}

func copyFields(src map[string]any) map[string]any {
	dst := make(map[string]any, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func logCallerFilePath(loggingFilePath string) string {
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l *defaultLogger) Log(level Level, v ...any) {
	l.logf(level, "", v...)
}

func (l *defaultLogger) Logf(level Level, format string, v ...any) {
	l.logf(level, format, v...)
}

func (l *defaultLogger) logf(level Level, format string, v ...any) {
	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerFilePath(file), line)
	}
	metadataMap := make(map[string]string)
	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		metadataMap[k] = fmt.Sprintf("%v", v)
	}
	sort.Strings(keys)
	metadata := ""
	for i, k := range keys {
		if i == 0 {
			metadata += fmt.Sprintf("%v: %v", k, fields[k])
		} else {
			metadata += fmt.Sprintf(" , %v: %v", k, fields[k])
		}
	}

	message := ""
	if format == "" {
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(format, v...)
	}

	t := time.Now().Format("2006-01-02 15:04:05")
	var name, logStr string
	if l.opts.Name != "" {
		name = "[" + l.opts.Name + "]"
	}
	if name == "" {
		logStr = fmt.Sprintf("%s: %v, %s \n", t, message, metadata)
	} else {
		logStr = fmt.Sprintf("%s %s: %v, %s\n", name, t, message, metadata)
	}

	if _, err := l.opts.Out.Write([]byte(logStr)); err != nil {
		log.Printf("log write error: %v \n", err)
	}
}

func NewDefaultLogger(opts ...OptionFunc) ILogger {
	options := Options{
		Level:           InfoLevel,
		Fields:          make(map[string]any),
		Out:             os.Stderr,
		CallerSkipCount: 3,
		Context:         context.Background(),
		Name:            "default",
	}
	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		panic(err)
	}
	return l
}
