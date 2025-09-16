package slogrus

import (
	"log/slog"

	sloglogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
)

func convertLevel(level logrus.Level) slog.Level {
	switch level {
	case logrus.TraceLevel:
		return slog.LevelDebug
	case logrus.DebugLevel:
		return slog.LevelDebug
	case logrus.InfoLevel:
		return slog.LevelInfo
	case logrus.WarnLevel:
		return slog.LevelWarn
	case logrus.ErrorLevel:
		return slog.LevelError
	case logrus.FatalLevel:
		return slog.LevelError
	case logrus.PanicLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func Convert(logE *logrus.Entry) *slog.Logger {
	data := make([]any, 0, len(logE.Data)*2)
	for k, v := range logE.Data {
		data = append(data, k, v)
	}
	return slog.New(sloglogrus.Option{
		Level:  convertLevel(logE.Logger.Level),
		Logger: logE.Logger,
	}.NewLogrusHandler()).With(data...)
}
