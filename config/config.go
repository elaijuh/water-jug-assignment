package config

type Properties struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"8080"`
}
