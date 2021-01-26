package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	MySql *MySqlConfiguration
}

type MySqlConfiguration struct {
	Database string
	Username string
	Password string
	Host     string
	Port     int
}

func Initialize() *Configuration {
	// Initialize Viper
	setUpViper()
	configuration := &Configuration{}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("unable to decode configuration, %v", err)
	}
	// fmt.Print(configuration)
	return configuration
}

// func initMySQL("mysql")
func setUpViper() {

	// Set the file name of the configurations file
	viper.SetConfigName("config.toml")

	// Set config file type
	viper.SetConfigType("toml")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to read configuration file: %s", err))
	}
}
