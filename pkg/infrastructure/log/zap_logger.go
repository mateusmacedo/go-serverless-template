package log

import (
	"go.uber.org/zap"

	"go-sls-template/pkg/application"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (application.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	return &ZapLogger{logger: sugar}, nil
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Error(msg string, err error, keysAndValues ...interface{}) {
	z.logger.Errorw(msg, append(keysAndValues, "error", err)...)
}
