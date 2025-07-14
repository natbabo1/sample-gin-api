package config

import (
	"errors"
	"fmt"
	"os"
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

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // sane default
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()      // ENV vars override file
	v.SetEnvPrefix("APP") // e.g. APP_DB_DSN

	if err := readExactFile(v, "./config.base.yaml"); err != nil {
		return nil, fmt.Errorf("base config: %w", err)
	}

	// 2. optionally merge config.<env>.yaml (may not exist)
	if err := mergeIfExists(v, fmt.Sprintf("./config.%s.yaml", env)); err != nil {
		return nil, err
	}
	var c Config
	return &c, v.Unmarshal(&c)
}

func readExactFile(v *viper.Viper, path string) error {
	v.SetConfigFile(path)
	return v.ReadInConfig() // returns error if file missing or bad YAML
}

func mergeIfExists(v *viper.Viper, path string) error {
	v.SetConfigFile(path)
	err := v.MergeInConfig() // returns ErrNotExist if file absent
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}
