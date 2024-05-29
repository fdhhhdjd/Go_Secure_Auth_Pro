package configs

import (
	"os"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	Gmail    GmailConfig
}

type ServerConfig struct {
	Host         string
	Port         string
	PortFrontend string
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

type GmailConfig struct {
	Host     string
	Port     string
	Password string
	Service  string
	Mail     string
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
