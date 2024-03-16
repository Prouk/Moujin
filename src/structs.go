package src

import "github.com/go-chi/chi/v5"

type Moujin struct {
	C   *Config
	R   *chi.Mux
	Dir string
}

type Config struct {
	Port string `yaml:"port"`
}
