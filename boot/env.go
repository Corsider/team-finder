package boot

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           int    `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"POSTGRES_USER"`
	DBPass           string `mapstructure:"POSTGRES_PASSWORD"`
	DBName           string `mapstructure:"POSTGRES_DB"`
	TokenSecret      string `mapstructure:"TOKEN_SECRET"`
	TokenTimeoutHour int    `mapstructure:"TOKEN_TIMEOUT_HOUR"`
	Timeout          int    `mapstructure:"TIMEOUT"`
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
