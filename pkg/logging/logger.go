package logging

import (
	"log"

	"go.uber.org/zap"
)

// NewLogger ... returns new zap logger
func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("couldn't build logger!!")
	}
	return logger
}
