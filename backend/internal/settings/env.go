package settings

import "github.com/caarlos0/env/v11"

type Environment struct {
	AppHost string `env:"APP_HOST,required"`
	AppPort int    `env:"APP_PORT" envDefault:"3000"`
	DatabaseEnvironment
}

type DatabaseEnvironment struct {
	DatabaseHost     string `env:"DATABASE_HOST,required"`
	DatabasePort     int    `env:"DATABASE_PORT,required"`
	DatabaseUser     string `env:"DATABASE_USER,required"`
	DatabaseName     string `env:"DATABASE_NAME,required"`
	DatabasePassword string `env:"DATABASE_PASSWORD,required"`
}

var envs Environment

func GetEnvs() *Environment {
	if (Environment{}) != envs {
		return &envs
	}

	if err := env.Parse(&envs); err != nil {
		panic(err)
	}
	return &envs
}
