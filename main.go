package main

import (
	"github.com/Prouk/Moujin/src"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Printf("starting moujin")
	m := new(src.Moujin)
	err := m.ReadConfigFromYaml("./config.yaml")
	if err != nil {
		log.Printf("error reading config file: %s", err)
		return
	}
	m.R = gin.Default()
	m.PopulateRoutes()
	err = m.R.Run(m.C.Port)
	if err != nil {
		log.Printf("error launching server: %s", err)
		return
	}
}
