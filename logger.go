package log

import (
	"os"
	"time"

	itbasisCoreUtils "github.com/itbasis/go-core-utils"
	itbasisDockerUtils "github.com/itbasis/go-docker-utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ConfigureDefaultContextLogger(forcePlainText bool) *zerolog.Logger {
	if forcePlainText {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			},
		)
	}

	log.Logger = log.With().
		Caller().
		Logger()
	zerolog.DefaultContextLogger = &log.Logger

	return zerolog.DefaultContextLogger
}

func ConfigureRootLogger(serviceName string) (*zerolog.Logger, error) {
	ConfigureDefaultContextLogger(false)

	config := Config{}
	if err := itbasisCoreUtils.ReadEnvConfig(&config, nil); err != nil {
		return nil, err
	}

	return ConfigureRootLoggerWithConfig(config, serviceName)
}

func ConfigureRootLoggerWithConfig(config Config, serviceName string) (*zerolog.Logger, error) {
	ConfigureDefaultContextLogger(!(config.LogForcePlainText || itbasisDockerUtils.IsRunningInDockerContainer()))

	zerolog.SetGlobalLevel(config.LogRootLevel.Level)

	zerolog.DefaultContextLogger.UpdateContext(
		func(c zerolog.Context) zerolog.Context {
			return c.Str(MdcServiceName, serviceName)
		},
	)

	return zerolog.DefaultContextLogger, nil
}
