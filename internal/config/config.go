package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"messenger-go/internal/logger"
	"messenger-go/internal/repository/postgres"
)

const (
	configPath = "configs"
	configName = "config"
)

type Config struct {
	Port     string
	DBConfig *postgres.Config
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.initConfig()

	return cfg
}

func (c *Config) initConfig() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading .env")
	}

	if err := viper.BindEnv("postgres_user"); err != nil {
		logger.Fatal("error binding db user env")
	}

	if err := viper.BindEnv("postgres_password"); err != nil {
		logger.Fatal("error binding db password env")
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	c.DBConfig = &postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("postgres_user"),
		Password: viper.GetString("postgres_password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	c.Port = viper.GetString("app.port")
}
