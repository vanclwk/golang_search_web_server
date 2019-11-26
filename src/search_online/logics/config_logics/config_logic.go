package config_logics

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func init() {
	fmt.Printf("Loading configuration logics...\n")
	GlobalConfig = initConfig()
	go dynamicConfig()
}

func initConfig() *viper.Viper {
	GlobalConfig := viper.New()
	GlobalConfig.SetConfigName("base")
	GlobalConfig.AddConfigPath("conf/")
	GlobalConfig.SetConfigType("yaml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to get the configuration.")
	}
	return GlobalConfig
}

func dynamicConfig() {
	GlobalConfig.WatchConfig()
	GlobalConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})
}