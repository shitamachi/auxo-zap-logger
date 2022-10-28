package main

import (
	"github.com/cuigh/auxo/log"
	"go.uber.org/zap/zapcore"
)

func ZapLevelToAuxoLevel(zl zapcore.Level) log.Level {
	var l log.Level
	switch zl {
	case zapcore.DebugLevel:
		l = log.LevelDebug
	case zapcore.InfoLevel:
		l = log.LevelInfo
	case zapcore.WarnLevel:
		l = log.LevelWarn
	case zapcore.ErrorLevel:
		l = log.LevelError
	case zapcore.DPanicLevel:
		l = log.LevelPanic
	case zapcore.PanicLevel:
		l = log.LevelPanic
	case zapcore.FatalLevel:
		l = log.LevelFatal
	case zapcore.InvalidLevel:
		l = log.LevelOff
	default:
		l = log.LevelOff
	}

	return l
}

func AuxoLevelToZapLevel(al log.Level) zapcore.Level {
	var l zapcore.Level
	switch al {
	case log.LevelDebug:
		l = zapcore.DebugLevel
	case log.LevelInfo:
		l = zapcore.InfoLevel
	case log.LevelWarn:
		l = zapcore.WarnLevel
	case log.LevelError:
		l = zapcore.ErrorLevel
	case log.LevelPanic:
		l = zapcore.PanicLevel
	case log.LevelFatal:
		l = zapcore.FatalLevel
	case log.LevelOff:
		l = zapcore.InvalidLevel
	default:
		l = zapcore.InvalidLevel
	}

	return l
}
