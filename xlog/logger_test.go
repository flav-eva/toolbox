package xlog

import "testing"

func TestDefault(t *testing.T) {
	DefaultLogger.Log(InfoLevel, "test world!")

	DefaultLogger.Infof("test info: %v", "hhahha")

	DefaultLogger.Info("default info", " hello default")

	Infof("default xlog for global: %v", "global")

}
