package zap

import (
	"fmt"
	"testing"

	"github.com/flav-eva/toolbox/xlog"
)

func TestLog(t *testing.T) {
	zlog := xlog.NewLogger(NewZapLogger())

	fmt.Println(zlog.Name())

	zlog.Infof("test zap: %v", "zap info")
}
