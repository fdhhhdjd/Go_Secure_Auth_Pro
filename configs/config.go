package configs

import (
	"os"

	constants "github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

type CacheConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	env := os.Getenv("ENV")
	if env == constants.DevEnvironment {
		viper.SetConfigName("config.dev")
	} else {
		viper.SetConfigName("config.prod")
	}
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
