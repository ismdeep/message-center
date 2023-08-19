package conf

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Server struct {
		Bind string `env:"MC_SERVER_BIND,default=0.0.0.0:9000"`
		Mode string `env:"MC_SERVER_MODE,default=release"`
	}
	Auth struct {
		Username string `env:"MC_AUTH_USERNAME,default=admin"`
		Password string `env:"MC_AUTH_PASSWORD,default=123456"`
	}
}

var ROOT config

func init() {
	if err := envconfig.Process(context.TODO(), &ROOT); err != nil {
		panic(err)
	}
}
