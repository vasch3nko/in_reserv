package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

// Config is the project configuration struct
// with struct tags that needs for parse fields
type Config struct {
	Server struct {
		Addr         string        `yaml:"addr"          json:"addr"          toml:"addr"          env:"ADDR"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"  json:"idle_timeout"  toml:"idle_timeout"  env:"IDLE_TIMEOUT"`
		ReadTimeout  time.Duration `yaml:"read_timeout"  json:"read_timeout"  toml:"read_timeout"  env:"READ_TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout" toml:"write_timeout" env:"WRITE_TIMEOUT"`
	} `yaml:"server" json:"server" toml:"server" env-prefix:"IN_RESERV_SERVER_"`

	DB struct {
		Host          string `yaml:"host"           json:"host"           toml:"host"           env:"HOST"`
		Port          string `yaml:"port"           json:"port"           toml:"port"           env:"PORT"`
		User          string `yaml:"user"           json:"user"           toml:"user"           env:"USER"`
		Password      string `yaml:"password"       json:"password"       toml:"password"       env:"PASSWORD"`
		Database      string `yaml:"database"       json:"database"       toml:"database"       env:"DATABASE"`
		SSLMode       string `yaml:"ssl_mode"       json:"ssl_mode"       toml:"ssl_mode"       env:"SSL_MODE"`
		MigrationPath string `yaml:"migration_path" json:"migration_path" toml:"migration_path" env:"MIGRATION_PATH"`
	} `yaml:"db" json:"db" toml:"db" env-prefix:"IN_RESERV_DB_"`
}

// NewConfig is the constructor of the Config struct
func NewConfig(path string) (*Config, error) {
	cfg := new(Config)
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
