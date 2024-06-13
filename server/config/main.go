package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	ENV     string `mapstructure:"APP_ENV"`
	XeroAPI string `mapstructure:"XERO_API"`
}

var config Config

func Init() {
	var err error
	config, err = load(".")
	if err != nil {
		config, err = loadEnv()
	}
	if err != nil {
		logrus.Fatal("Could not load config: ", err)
	}
}

func Get() Config {
	return config
}

func load(path string) (conf Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	return
}

func loadEnv() (conf Config, err error) {
	viper.AutomaticEnv()

	conf = Config{
		AppPort: viper.GetString("app_port"),
		ENV:     viper.GetString("app_env"),
		XeroAPI: viper.GetString("xero_api"),
	}

	err = nil

	return
}
