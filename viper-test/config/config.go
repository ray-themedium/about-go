package config

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	Count  int    `mapstructure:"COUNT"`
}

func Load(path string, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
