package xlog

import (
	"testing"

	"go.uber.org/zap"
)

func TestOpenZapLogger(t *testing.T) {
	logger, err := OpenZapLogger(&ZapCFG{
		Development: true,
		Debug:       true,
		Fields: &Fields{
			App: "api-test",
		},
		Lumberjack: &Lumberjack{
			LogPath:    "zap_test.log",
			MaxSize:    1,
			MaxBackups: 20,
			MaxAge:     2,
		},
	})
	if err != nil {
		t.Error(err)
	}
	logger.Sugar().Infof("hello: %s", "world!")
}

func BenchmarkOpenZapLogger(b *testing.B) {
	logger, err := OpenZapLogger(&ZapCFG{
		Development: true,
		Debug:       true,
		Fields: &Fields{
			App: "api-test",
		},
		Lumberjack: &Lumberjack{
			LogPath:    "zap_test.log",
			MaxSize:    1,
			MaxBackups: 20,
			MaxAge:     2,
		},
	})
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark zapLogger", zap.Any("index", i))
	}
}
