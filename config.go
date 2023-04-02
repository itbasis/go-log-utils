package log

type Config struct {
	LogRootLevel      EnvLevel `env:"LOG_ROOT_LEVEL" envDefault:"INFO"`
	LogForcePlainText bool     `env:"LOG_FORCE_PLAIN_TEXT"`
}
