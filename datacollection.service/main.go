package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// Load the config file using viper the
	// config file holds you access keys
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

}
