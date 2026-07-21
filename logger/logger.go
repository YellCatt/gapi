package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Init 初始化日志组件。
// path  日志文件存放目录；level 控制台最低输出级别（debug/info/warn/error）。
func Init(path string, level string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create log directory %s: %w", path, err)
	}

	lvl := zapcore.InfoLevel
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		return fmt.Errorf("invalid log level %q: %w", level, err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	infoFile, err := openLogFile(path, "info.log")
	if err != nil {
		return err
	}
	warnFile, err := openLogFile(path, "warn.log")
	if err != nil {
		return err
	}
	errorFile, err := openLogFile(path, "error.log")
	if err != nil {
		return err
	}

	infoCore := zapcore.NewCore(encoder, zapcore.AddSync(infoFile), exactLevel(lvl, zapcore.InfoLevel))
	warnCore := zapcore.NewCore(encoder, zapcore.AddSync(warnFile), exactLevel(lvl, zapcore.WarnLevel))
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorFile), atLeastLevel(zapcore.ErrorLevel))
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atLeastLevel(lvl))

	core := zapcore.NewTee(infoCore, warnCore, errorCore, consoleCore)

	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(log)
	return nil
}

func openLogFile(path, name string) (*os.File, error) {
	file, err := os.OpenFile(filepath.Join(path, name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %s: %w", name, err)
	}
	return file, nil
}

// exactLevel 只输出指定级别，且不低于用户配置的最低级别。
func exactLevel(min, target zapcore.Level) zapcore.LevelEnablerFunc {
	return func(l zapcore.Level) bool {
		return l >= min && l == target
	}
}

// atLeastLevel 输出该级别及以上日志。
func atLeastLevel(min zapcore.Level) zapcore.LevelEnablerFunc {
	return func(l zapcore.Level) bool {
		return l >= min
	}
}

func Info(msg string, fields ...zap.Field) {
	if log == nil {
		return
	}
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	if log == nil {
		return
	}
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if log == nil {
		return
	}
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	if log == nil {
		return
	}
	log.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	if log == nil {
		return
	}
	log.Fatal(msg, fields...)
}

func Sync() {
	if log == nil {
		return
	}
	_ = log.Sync()
}
