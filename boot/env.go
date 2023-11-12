package boot

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        int    `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	TokenSecret   string `mapstructure:"TOKEN_SECRET"`
	Timeout       int    `mapstructure:"TIMEOUT"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("No .env file in project root", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error while reading .env file", err)
	}

	log.Println("Config loaded")
	return &env
}
