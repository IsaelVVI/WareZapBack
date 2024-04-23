package main

import (
	"github.com/IsaelVVI/warezapback.git/config"
	"github.com/IsaelVVI/warezapback.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize Configs app
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialize error: %v", err)
		return
	}

	// Inicializando o package Router e API
	router.Initialize()
}
