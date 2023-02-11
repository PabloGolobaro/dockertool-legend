package logger

import (
	"go.uber.org/zap"
	"os"
)

func NewLogger() (*zap.SugaredLogger, error) {
	var err error
	var log *zap.Logger
	if os.Getenv("MODE") == "dev" {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	sugar := log.Sugar()
	return sugar, nil
}
