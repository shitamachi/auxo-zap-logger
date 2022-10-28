package auxozap

import (
	"github.com/cuigh/auxo/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapio"
)

type ZapLogger struct {
	name   string
	lvl    zapcore.Level
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	writer *zapio.Writer
}

func NewZapLogger(name string, lvl zapcore.Level, logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		name:   name,
		lvl:    lvl,
		logger: logger,
		sugar:  logger.Sugar(),
		writer: &zapio.Writer{
			Log:   logger,
			Level: lvl,
		},
	}
}

func (z ZapLogger) Write(p []byte) (n int, err error) {
	return z.writer.Write(p)
}

func (z ZapLogger) Name() string {
	return z.name
}

func (z ZapLogger) Level() log.Level {
	return ZapLevelToAuxoLevel(z.lvl)
}

func (z ZapLogger) SetLevel(lvl log.Level) {
	AuxoLevelToZapLevel(lvl)
}

func (z ZapLogger) IsEnabled(lvl log.Level) bool {
	return ZapLevelToAuxoLevel(z.lvl) >= lvl
}

func (z ZapLogger) WithField(key string, value interface{}) log.Entry {
	l := z.logger.With(zap.Any(key, value))
	return ZapLogger{
		name:   z.name,
		lvl:    z.lvl,
		logger: l,
		sugar:  l.Sugar(),
		writer: &zapio.Writer{Log: l, Level: z.lvl},
	}
}

func (z ZapLogger) WithFields(fields map[string]interface{}) log.Entry {
	zapFields := make([]zap.Field, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	l := z.logger.With(zapFields...)
	return ZapLogger{
		name:   z.name,
		lvl:    z.lvl,
		logger: l,
		sugar:  l.Sugar(),
		writer: &zapio.Writer{Log: l, Level: z.lvl},
	}
}

func (z ZapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}

func (z ZapLogger) Debugf(format string, args ...interface{}) {
	z.sugar.Debugf(format, args...)
}

func (z ZapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}

func (z ZapLogger) Infof(format string, args ...interface{}) {
	z.sugar.Infof(format, args...)
}

func (z ZapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}

func (z ZapLogger) Warnf(format string, args ...interface{}) {
	z.sugar.Warnf(format, args...)
}

func (z ZapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}

func (z ZapLogger) Errorf(format string, args ...interface{}) {
	z.sugar.Errorf(format, args...)
}

func (z ZapLogger) Panic(args ...interface{}) {
	z.sugar.Panic(args...)
}

func (z ZapLogger) Panicf(format string, args ...interface{}) {
	z.sugar.Panicf(format, args...)
}

func (z ZapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args...)
}

func (z ZapLogger) Fatalf(format string, args ...interface{}) {
	z.sugar.Fatalf(format, args...)
}
