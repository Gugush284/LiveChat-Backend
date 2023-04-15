package jsonServer

// Config ...
type jsonConfig struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	SessionKey string `toml:"session_key"`
}

// New Config
func NewConfig() *jsonConfig {
	return &jsonConfig{
		BindAddr:   ":8080",
		LogLevel:   "debug",
		SessionKey: "NoKey",
	}
}
