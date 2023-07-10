package logrus

import (
	"testing"

	"github.com/flav-eva/toolbox/xlog"
)

func TestName(t *testing.T) {
	l := xlog.NewLogger(NewLogrusLogger())

	if l.Name() != "logrus" {
		t.Errorf("error: name expected 'logrus' actual: %s", l.Name())
	}

	t.Logf("testing xlog name: %s", l.Name())

	l.Infof("test logrus: %v", "logrus info")
}
