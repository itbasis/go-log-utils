package log

import "go.uber.org/zap/zapcore"

type Config struct {
	LogRootLevel      zapcore.Level `env:"LOG_ROOT_LEVEL" envDefault:"info"`
	LogForcePlainText bool          `env:"LOG_FORCE_PLAIN_TEXT"`
}
