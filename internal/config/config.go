package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port string `mapstructure:"port"`
    } `mapstructure:"server"`
    Integrations struct {
        Slack struct {
            WebhookURL string `mapstructure:"webhook_url"`
        } `mapstructure:"slack"`
    } `mapstructure:"integrations"`
}

func LoadConfig(path string) (config *Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    err = viper.ReadInConfig()
    if err != nil {
        return
    }
    err = viper.Unmarshal(&config)
    return
}