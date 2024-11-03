package config

import "github.com/spf13/viper"
type Config struct {
	DbHost     string `mapstructure:"POSTGRES_HOST"`
	DbUsername string `mapstructure:"POSTGRES_USER"`
	DbPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DbName     string `mapstructure:"POSTGRES_DB"`
	DbPort     string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
    viper.SetConfigType("env")
    viper.SetConfigName("app")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)

    return
}
