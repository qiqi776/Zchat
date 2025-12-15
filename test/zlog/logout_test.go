package zlog

import (
	"Zchat/pkg/zlog"
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	zlog.Info("this is a info", zap.String("name", "apylee"))
}