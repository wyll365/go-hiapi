package hiapi_basic_config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func ReadConfig() *ApplicationConfig {

	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	var config ApplicationConfig
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct： %s", err))
	}
	if len(config.Environment) != 0 {
		return readEnv(config)
	}
	return &config
}

func readEnv(base ApplicationConfig) *ApplicationConfig {
	viper.SetConfigFile("application-" + strings.ToLower(base.Environment) + ".yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error environment config file: %s", err))
	}

	var config ApplicationConfig
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into environment  struct： %s", err))
	}

	return &config
}
