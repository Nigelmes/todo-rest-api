package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			text := "Error reading configuration"
			help, _ := cleanenv.GetDescription(instance, &text)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}