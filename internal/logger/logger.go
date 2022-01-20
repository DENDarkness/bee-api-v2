package logger

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	// TODO: при надобности можно переделать на кастомный конфиг.
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"/var/log/bee-api/bee.log",
	}
	// logger, err := zap.NewProduction()
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	return logger
}
