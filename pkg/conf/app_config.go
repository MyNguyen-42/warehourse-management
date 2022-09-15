package conf

import (
	"github.com/spf13/viper"
	"os"
)

type (
	Database struct {
		Hosts      []string `mapstructure:"hosts,omitempty"`
		Username   string   `mapstructure:"username,omitempty"`
		Password   string   `mapstructure:"password,omitempty"`
		AuthSource string   `mapstructure:"authSource,omitempty"`
	}
	Config struct {
		Env           string   `mapstructure:"env,omitempty"`
		Port          string   `mapstructure:"port,omitempty"`
		Database      Database `mapstructure:"database,omitempty"`
		ProductDbName string   `mapstructure:"productDbName,omitempty"`
		StoreDbName   string   `mapstructure:"storeDbName,omitempty"`
	}
)

var config *Config

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("conf-" + os.Getenv("APP_ENV"))
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func GetConfig() Config {
	return *config
}

func LoadConfigFiles(path string) (config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("conf-" + os.Getenv("APP_ENV"))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
