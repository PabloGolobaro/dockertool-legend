package logger

import "go.uber.org/zap"

func NewLogger() (*zap.SugaredLogger, error) {
	development, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	sugar := development.Sugar()
	return sugar, nil
}
