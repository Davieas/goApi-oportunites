package main

import (
	"github.com/Davieas/goapi-oportunites/config"
	"github.com/Davieas/goapi-oportunites/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Init()
}
