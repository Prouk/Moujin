package src

import (
	"github.com/gin-gonic/gin"
)

type Moujin struct {
	C *Config
	R *gin.Engine
}

type Config struct {
	Port string `yaml:"port"`
}
