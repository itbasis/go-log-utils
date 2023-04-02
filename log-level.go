package log

import (
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type EnvLevel struct {
	Level zerolog.Level
}

func (receiver *EnvLevel) UnmarshalText(text []byte) error {
	v := strings.ToLower(string(text))

	var level zerolog.Level

	switch v {
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "debug":
		level = zerolog.DebugLevel
	case "trace":
		level = zerolog.TraceLevel
	default:
		log.Warn().Msgf("Unknown logging level: %s", text)

		level = zerolog.InfoLevel
	}

	*receiver = EnvLevel{Level: level}

	return nil
}
