package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string 
   Port      int 
   Host      string
   User      string
   DB        string
   Password  string 
}

func Load() Config {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	v := viper.New()
	v.AutomaticEnv()

    config.HTTPPort = v.GetString("HTTP_PORT")
    config.DB = v.GetString("POSTGRES_DB")
	config.Host = v.GetString("POSTGRES_HOST")
	config.User = v.GetString("POSTGRES_USER")
	config.Port = v.GetInt("POSTGRES_PORT")
	config.Password = v.GetString("POSTGRES_PASSWORD")

  return config
}