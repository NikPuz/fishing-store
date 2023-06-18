package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" validate:"required"`

	Server `yaml:"server" validate:"required"`
}

type Server struct {
	Port         int           `yaml:"port" validate:"required"`
	WriteTimeout time.Duration `yaml:"writeTimeout" validate:"required"`
	ReadTimeout  time.Duration `yaml:"readTimeout" validate:"required"`
	IdleTimeout  time.Duration `yaml:"idleTimeout" validate:"required"`
}

func NewConfig() *Config {
	cfg := new(Config)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	if err := validator.New().Struct(cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
