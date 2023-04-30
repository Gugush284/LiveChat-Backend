package server

// Config ...
type config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// New Config
func NewConfig() *config {
	return &config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
