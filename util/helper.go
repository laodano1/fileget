package util

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func init() {
	//Logger, _ = zap.NewProduction()
	Logger, _ = zap.NewDevelopment()
}
