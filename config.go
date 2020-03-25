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
		Port           string `yaml:"PORT"`
		UserHost       string `yaml:"USERHOST"`
		FriendHost     string `yaml:"FRIENDHOST"`
		RecipeHost     string `yaml:"RECIPEHOST"`
		Friendname     string `yaml:"FRIENDNAME"`
		Username       string `yaml:"USERNAME"`
		RecipeName     string `yaml:"RECIPENAME"`
		UserUsername   string `yaml:"USERUSERNAME"`
		UserPassword   string `yaml:"USERPASSWORD"`
		FriendUsername string `yaml:"FRIENDUSERNAME"`
		FriendPassword string `yaml:"FRIENDPASSWORD"`
		RecipeUsername string `yaml:"RECIPEUSERNAME"`
		RecipePassword string `yaml:"RECIPEPASSWORD"`
		HerokuUser     string `yaml:"HEROKUUSER"`
		HerokuFriend   string `yaml:"HEROKUFRIEND"`
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
