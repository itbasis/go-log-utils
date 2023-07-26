package log

import (
	"context"
	"log"

	itbasisCoreUtils "github.com/itbasis/go-core-utils/v2"
	itbasisDockerUtils "github.com/itbasis/go-docker-utils/v2"
	"github.com/juju/zaputil/zapctx"
	"go.uber.org/zap"
)

func ConfigureDefaultContextLogger(forcePlainText bool, zapConfig zap.Config) *zap.Logger {
	if forcePlainText {
		zapConfig.Encoding = "console"
	}

	return ConfigureDefaultContextLoggerCustomConfig(zapConfig)
}

func ConfigureDefaultContextLoggerCustomConfig(config zap.Config) *zap.Logger {
	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		log.Panic(err)
	}

	zapctx.Default = logger

	return zapctx.Default
}

func ConfigureRootLogger(ctx context.Context, serviceName string, zapConfig zap.Config) (*zap.Logger, error) {
	ConfigureDefaultContextLogger(false, zapConfig)

	config := Config{}
	if err := itbasisCoreUtils.ReadEnvConfig(ctx, &config, nil); err != nil {
		return nil, err
	}

	return ConfigureRootLoggerWithConfig(serviceName, zapConfig, config)
}

func ConfigureRootLoggerWithConfig(serviceName string, zapConfig zap.Config, config Config) (*zap.Logger, error) {
	ConfigureDefaultContextLogger(
		!(config.LogForcePlainText || itbasisDockerUtils.IsRunningInDockerContainer()),
		zapConfig,
	)

	zapctx.LogLevel.SetLevel(config.LogRootLevel)
	zapctx.Default = zapctx.Default.With(zap.String(MdcServiceName, serviceName))

	return zapctx.Default, nil
}
