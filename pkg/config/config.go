package config

import "github.com/AnnaKuvarina/pet-go-app/pkg/stores"

type Config struct {
	Port         uint
	PostgreStore *stores.PostgreConfig
	MongoStore   *stores.MongoConfig
}

var AppConfig = &Config{}
