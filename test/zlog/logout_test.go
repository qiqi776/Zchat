package zlog

import (
	"go.uber.org/zap"
	"testing"
	"Zchat/pkg/zlog"
)

func TestInfo(t *testing.T) {
	zlog.Info("this is a info", zap.String("name", "apylee"))
}