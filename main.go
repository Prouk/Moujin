package main

import (
	"github.com/Prouk/Moujin/src"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("starting moujin")
	m := new(src.Moujin)
	err := m.ReadConfigFromYaml("./config.yaml")
	if err != nil {
		log.Printf("error reading config file: %s", err)
		return
	}
	m.R = chi.NewRouter()
	m.Dir, err = os.Getwd()
	if err != nil {
		return
	}
	m.PopulateRoutes()
	err = http.ListenAndServe(m.C.Port, m.R)
	if err != nil {
		return
	}
}
