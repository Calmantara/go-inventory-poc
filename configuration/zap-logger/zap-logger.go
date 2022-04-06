package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(z *zap.Config)

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func NewWrappedZapLogger(ops ...Option) *zap.SugaredLogger {
	// logger setup
	cfg := zap.Config{
		Encoding:          "console",
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:       []string{"stderr"},
		Development:       false,
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    customLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	//iterate all option function
	for _, v := range ops {
		v(&cfg)
	}
	// initiate logger and sugar logger
	logger, _ := cfg.Build()
	logger.Info("Initialization zap logger. . .")
	return logger.Sugar()
}

func WithTimeKey(tk string) Option {
	return func(z *zap.Config) { z.EncoderConfig.TimeKey = tk }
}
func WithTimeFormat(encodeTime zapcore.TimeEncoder) Option {
	return func(z *zap.Config) { z.EncoderConfig.EncodeTime = encodeTime }
}
