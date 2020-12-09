package conf

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

type ServiceConfiguration struct {
	Host       string `env:"INTKEY_HOST" envDefault:"0.0.0.0"`
	Port       uint16 `env:"INTKEY_PORT" envDefault:"8080"`
	TimeFormat string `env:"INTKEY_TIME_FORMAT" envDefault:"default"`
}

var instantiated *ServiceConfiguration
var once sync.Once

func ServiceConfigurationIns() *ServiceConfiguration {
	once.Do(func() {
		instantiated = &ServiceConfiguration{}
		if err := env.Parse(instantiated); err != nil {
			log.Fatal().Msg(err.Error())
			os.Exit(1)
		}
	})
	return instantiated
}
