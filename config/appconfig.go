package config

import (
	"fmt"
	"log"
	"postgre-basic/database/postgre"

	"github.com/spf13/viper"
)

type Application struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

type AppConfiguration struct {
	Application Application
	Database    *postgre.PostgreConfig
}

var conf *AppConfiguration

func GetConfig() *AppConfiguration {

	conf = InitConfiguration()

	return conf
}

func InitDevConfiguration() *AppConfiguration {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Error reading config", err.Error())
		panic(err)
	}

	configApp := &AppConfiguration{}

	err = viper.Unmarshal(&configApp)

	if err != nil {
		log.Println("Errpr unmarshal config", err.Error())
		panic(err)
	}

	return configApp
}

func InitConfiguration() *AppConfiguration {
	viper.AddConfigPath("./")
	viper.SetConfigName("config-local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config", err.Error())
		panic(err)
	}

	configApp := &AppConfiguration{}

	err = viper.Unmarshal(&configApp)

	if err != nil {
		fmt.Println("Errpr unmarshal config", err.Error())
		panic(err)
	}

	return configApp
}
