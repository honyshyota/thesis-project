package app

import (
	configuration "main/configs"
	"main/internal/router"
)

func Run() {
	cfg := configuration.CheckCfg() // Launching a common config
	router.Router(cfg)
}
