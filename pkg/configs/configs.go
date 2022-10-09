package configs

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

type Config struct {
	Server   *ServerConfig   `json:"server" mapstructure:"server"`
	DataBase *DataBaseConfig `json:"database" mapstructure:"database"`
}

type ServerConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Port    int    `json:"port" mapstructure:"port"`
}

type DataBaseConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	DB       string `json:"database" mapstructure:"database"`
	Migrate  bool   `json:"migrate" mapstructure:"migrate"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

func init() {
	pflag.StringVarP(&cfgFile, "config", "c", cfgFile, "Read configuration from specified `FILE`")
}

func ReadConfig() (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("configs")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read configuration file(%s): %v", cfgFile, err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration file(%s): %v", cfgFile, err)
	}
	return &cfg, nil
}
