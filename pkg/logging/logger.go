package logging

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("couldn't build logger!!")
	}
	return logger
}
