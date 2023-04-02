package log_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/itbasis/go-log-utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

const (
	envLogRootLevelName = "LOG_ROOT_LEVEL"
)

type EnvLevelTestSuite struct {
	suite.Suite

	oldEnvRootLevel  string
	oldEnvForcePlain string
}

func TestEnvLevelSuite(t *testing.T) {
	suite.Run(t, &EnvLevelTestSuite{})
}

func (s *EnvLevelTestSuite) SetupSuite() {
	s.oldEnvRootLevel = os.Getenv(envLogRootLevelName)
	s.oldEnvForcePlain = os.Getenv("LOG_FORCE_PLAIN_TEXT")
}

func (s *EnvLevelTestSuite) TearDownSuite() {
	s.NoError(os.Setenv(envLogRootLevelName, s.oldEnvRootLevel))
	s.NoError(os.Setenv("LOG_FORCE_PLAIN_TEXT", s.oldEnvForcePlain))
}

func (s *EnvLevelTestSuite) TestReadEnvironment() {
	tests := []struct {
		envRootLevel string
		expect       log.Config
	}{
		{envRootLevel: "info", expect: log.Config{LogRootLevel: log.EnvLevel{Level: zerolog.InfoLevel}}},
		{envRootLevel: "debug", expect: log.Config{LogRootLevel: log.EnvLevel{Level: zerolog.DebugLevel}}},
		{envRootLevel: "trace", expect: log.Config{LogRootLevel: log.EnvLevel{Level: zerolog.TraceLevel}}},
		{envRootLevel: "Trace", expect: log.Config{LogRootLevel: log.EnvLevel{Level: zerolog.TraceLevel}}},
		{envRootLevel: "TRACE", expect: log.Config{LogRootLevel: log.EnvLevel{Level: zerolog.TraceLevel}}},
	}

	for i, test := range tests {
		s.Run(
			fmt.Sprintf("%d :: %s", i, test.envRootLevel), func() {
				s.NoError(os.Setenv(envLogRootLevelName, test.envRootLevel))

				config := &log.Config{}
				s.NoError(env.Parse(config))
			},
		)
	}
}

// func TestParseEnvLoggerEnv(t *testing.T) {
// 	tests := []struct {
// 		level  string
// 		expect zerolog.Level
// 	}{
// 		{"info", zerolog.InfoLevel},
// 		{"Info", zerolog.InfoLevel},
// 		{"inFo", zerolog.InfoLevel},
// 		{"warn", zerolog.WarnLevel},
// 		{"error", zerolog.ErrorLevel},
// 		{"debug", zerolog.DebugLevel},
// 		{"trace", zerolog.TraceLevel},
// 		{"", zerolog.InfoLevel},
// 		{"-", zerolog.InfoLevel},
// 	}
// 	for _, test := range tests {
// 		t.Run(
// 			test.level, func(t *testing.T) {
// 				assert.Equal(t, test.expect, logUtils.ParseEnvLoggerEnv(test.level))
// 			},
// 		)
// 	}
// }
