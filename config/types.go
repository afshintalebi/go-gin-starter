package config

// Configurations exported
type Configurations struct {
	App    AppConfigurations
	Sentry SentryConfigurations
}

// AppConfigurations exported
type AppConfigurations struct {
	Name string
	Env  string `mapstructure:"GIN_MODE"`
}

// AppConfigurations exported
type SentryConfigurations struct {
	DSN string `mapstructure:"SENTRY_DSN"`
}
