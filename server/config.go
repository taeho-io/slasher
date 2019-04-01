package server

type Config interface {
	Settings() Settings
}

type DefaultConfig struct {
	Config

	settings Settings
}

func NewConfig(settings Settings) (Config, error) {
	return &DefaultConfig{
		settings: settings,
	}, nil
}

func MockConfig() Config {
	cfg, _ := NewConfig(MockSettings())
	return cfg
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}
