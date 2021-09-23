package config

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/spf13/viper"
    "time"
)

type Configuration struct {
    Client              domains.Clients
    LoginEndpoint       string        `mapstructure:"login_endpoint"`
    AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
    AccessLogEnabled    bool          `mapstructure:"access_log_enabled"`
}

func Load() (*Configuration, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        return nil, errors.New("unable read config")
    }

    var config Configuration
    err = viper.Unmarshal(&config)
    if err != nil {
        return nil, errors.New("unable read config")
    }
    return &config, nil
}
