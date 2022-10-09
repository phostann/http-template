package log

import "go.uber.org/zap"

type Logger = zap.SugaredLogger

func NewLogger() *Logger {
	log, _ := zap.NewDevelopment()
	return log.Sugar()
}
