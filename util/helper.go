package util

import (
	"go.uber.org/zap"
)

var (
	ZLogger *zap.SugaredLogger
)

func init() {
	//Logger, _ = zap.NewProduction()

	zlog, _ := zap.NewDevelopment()
	ZLogger = zlog.Sugar()

}
