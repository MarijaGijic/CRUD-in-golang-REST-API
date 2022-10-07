package main

import (
	"log"

	"github.com/spf13/viper"
)

//struct that will contain the allowed configurations
type Config struct {
	Port             string `mapstructure:"port"`              //port num
	ConnectionString string `mapstructure:"connection_string"` //mysql connection string
}

var AppConfig *Config //define the AppConfig variable that will be accessed by other files and packages
//withing the application code
//use Viper to load configurations from congig.json file(will create) and
//assign its values to the AppConfig variable
func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
