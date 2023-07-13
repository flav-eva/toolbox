package zap_filter

import (
	"encoding/json"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const filterJsonEncoding = "filterjson"

type filterJsonEncoder struct {
	*zapcore.EncoderConfig
	zapcore.Encoder
	filter *filter
}

type filter struct {
	rules []FilterRule
}

type FilterRule interface {
	Filter(data map[string]interface{})
}

// NewFilterJSONEncoder ...
func (f *filter) NewFilterJSONEncoder(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
	return &filterJsonEncoder{
		EncoderConfig: &cfg,
		Encoder:       zapcore.NewJSONEncoder(cfg),
		filter:        f,
	}, nil
}

func (f *filter) handle(logBuffer *buffer.Buffer) *buffer.Buffer {
	if len(f.rules) == 0 {
		return logBuffer
	}
	jsonMap := make(map[string]interface{})
	json.Unmarshal(logBuffer.Bytes(), &jsonMap)
	for _, rule := range f.rules {
		rule.Filter(jsonMap)
	}
	data, _ := json.Marshal(jsonMap)
	logBuffer.Reset()
	logBuffer.Write(data)
	logBuffer.AppendString(zapcore.DefaultLineEnding)
	return logBuffer
}

func (enc *filterJsonEncoder) Clone() zapcore.Encoder {
	return &filterJsonEncoder{
		EncoderConfig: enc.EncoderConfig,
		Encoder:       enc.Encoder.Clone(),
		filter:        enc.filter,
	}
}

// EncodeEntry partially implements the zapcore.Encoder interface.
func (enc *filterJsonEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(ent, fields)
	if err != nil {
		return nil, err
	}

	buf = enc.filter.handle(buf)
	return buf, nil
}
