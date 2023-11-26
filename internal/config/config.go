package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	AppHost     string `env:"APP_HOST" env-required:"true"`
	AppPort     string `env:"APP_PORT" env-required:"true"`
	AppEnv      string `env:"APP_ENV" env-required:"true"`
	PG_Host     string `env:"PG_HOST" env-required:"true"`
	PG_Port     string `env:"PG_PORT" env-required:"true"`
	PG_User     string `env:"PG_USER" env-required:"true"`
	PG_Password string `env:"PG_PASSWORD" env-required:"true"`
	PG_DB       string `env:"PG_DB" env-required:"true"`
}

func Load() Config {
	var cfg Config
	appErr := cleanenv.ReadConfig(".env", &cfg)
	if appErr != nil {
		errMessage := fmt.Errorf("Read env err: %v.", appErr)
		fmt.Println(errMessage)
		os.Exit(1)
	}
	return cfg

}
