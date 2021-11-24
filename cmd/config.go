package main

import (
	"github.com/AnnaKuvarina/pet-go-app/pkg/config"
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

const (
	envFileName = "ENV_FILE_NAME"
	postgreName = "POSTGRE_NAME"
	postgreHost = "POSTGRE_HOST"
	postgrePort = "POSTGRE_PORT"
	postgreUser = "POSTGRE_USER"
	postgrePass = "POSTGRE_PASSWORD"
	mongoURL    = "MONGO_URL"
)

const (
	configFileNameDefault  = "config"
	defaultPostgreHost     = "localhost"
	defaultPostgreName     = "postgre"
	defaultPostgrePort     = 5030
	defaultPostgreUser     = "postgre"
	defaultPostgrePassword = 1234
	defaultMongoURL        = ""
)

func LoadConfig() error {
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetDefault(envFileName, configFileNameDefault)
	configFileName := viper.GetString(envFileName)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info().Str("filename", configFileName).Msg("Config file not set")
		} else {
			return err
		}
	}

	// postgre
	viper.SetDefault(postgreHost, defaultPostgreHost)
	viper.SetDefault(postgrePort, defaultPostgrePort)
	viper.SetDefault(postgreName, defaultPostgreName)
	viper.SetDefault(postgreUser, defaultPostgreUser)
	viper.SetDefault(postgrePass, defaultPostgrePassword)

	// mongo
	viper.SetDefault(mongoURL, defaultMongoURL)

	config.AppConfig = &config.Config{
		PostgreStore: &stores.PostgreConfig{
			Host:               viper.GetString(postgreHost),
			DB:                 viper.GetString(postgreName),
			Port:               viper.GetInt(postgrePort),
			User:               viper.GetString(postgreUser),
			Password:           viper.GetString(postgrePass),
			MaxConnections:     10,
			MaxIdleConnections: 10,
			MaxRetries:         10,
			ConnMaxLifetime:    time.Duration(10),
			ConnSleepPeriod:    time.Duration(2),
		},
		MongoStore: &stores.MongoConfig{
			DBUrl: viper.GetString(mongoURL),
		},
	}

	return nil
}
