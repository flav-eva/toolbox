package xlog

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func OpenZapLogger(cfg *ZapCFG) (*zap.Logger, error) {
	var zapCfg zap.Config
	options := []zap.Option{
		zap.AddCallerSkip(1 + cfg.CallerSkip),
	}

	var level zapcore.Level
	if cfg.Development {
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		level = zapcore.DebugLevel
	} else {
		zapCfg = zap.NewProductionConfig()
		options = append(options, zap.Fields(
			zap.String("_APP_", cfg.Fields.App)),
		)
		level = zapcore.InfoLevel
	}
	zapCfg.EncoderConfig.LevelKey = "_LEVEL_"
	zapCfg.EncoderConfig.TimeKey = "_TS_"
	zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	zapCfg.EncoderConfig.NameKey = "_NAME_"
	zapCfg.EncoderConfig.MessageKey = "_MSG_"
	zapCfg.EncoderConfig.CallerKey = "_CALLER_"
	zapCfg.EncoderConfig.StacktraceKey = "_STACKTRACE_"

	if !cfg.Sample { // 覆盖默认的采样率
		zapCfg.Sampling = nil
	}
	if cfg.Debug {
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	zapCfg.InitialFields = cfg.Fields.ExtraFields
	zapCfg.DisableStacktrace = cfg.DisableStackTrace

	writeSyncer := zapcore.AddSync(os.Stdout)
	if cfg.Lumberjack != nil {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Lumberjack.LogPath,
			MaxSize:    cfg.Lumberjack.MaxSize,
			MaxAge:     cfg.Lumberjack.MaxAge,
			MaxBackups: cfg.Lumberjack.MaxBackups,
			Compress:   cfg.Lumberjack.Compress,
		}))
	}

	// TODO 可以实现一个全新的 encoder，然后进行替换

	enc := zapcore.NewJSONEncoder(zapCfg.EncoderConfig)
	logger := zap.New(zapcore.NewCore(enc, writeSyncer, level), options...)

	return logger, nil
}
