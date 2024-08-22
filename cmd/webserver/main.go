package main

import (
	"github.com/joserafaelSH/go-api/config/env"
	"github.com/joserafaelSH/go-api/config/logger"
	"github.com/joserafaelSH/go-api/internal/handler"
)

func main() {
	logger.CreateLog("WebServer")
	err := env.LoadEnv()
	if err != nil {
		logger.Log.Writter.Error().Msg("Error loading environment variables")
		return
	}

	handler.InitHttpHandler()
}
