package recipes

import (
	"log"

	"github.com/spf13/viper"
)

//Config configuration
type Config struct {
	Server struct {
		Port string `yaml:"PORT"`
		Host string `yaml:"HOST"`
	} `yaml:"SERVER"`

	Storage struct {
		Port         string `yaml:"PORT"`
		Host         string `yaml:"HOST"`
		Username     string `yaml:"USERNAME"`
		Name         string `yaml:"NAME"`
		Password     string `yaml:"PASSWORD"`
		HerokuUser   string `yaml:"HEROKUUSER"`
		HerokuFriend string `yaml:"HEROKUFRIEND"`
	} `yaml:"STORAGE"`
}

//Conf config
func Conf() Config {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to find config file, error: %s\n", err.Error())
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to unmarshal into struct, error: %s\n", err.Error())
	}
	return conf
}
