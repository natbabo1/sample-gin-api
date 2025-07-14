package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port         int           `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	}
	DB struct {
		DSN          string `mapstructure:"dsn"` // postgres://user:pass@host:5432/db?sslmode=disable
		MaxIdleConns int    `mapstructure:"max_idle"`
		MaxOpenConns int    `mapstructure:"max_open"`
	}
	Env string `mapstructure:"env"` // dev | prod | test
}

func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AutomaticEnv()      // ENV vars override file
	v.SetEnvPrefix("APP") // e.g. APP_DB_DSN
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	return &c, v.Unmarshal(&c)
}
