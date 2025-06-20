package util

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger(level string) {
	var err error
	cfg := zap.NewProductionConfig()

	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	}

	Logger, err = cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
