package main

import (
	"log/slog"

	"github.com/joserafaelSH/go-api/config/env"
	"github.com/joserafaelSH/go-api/config/logger"
	"github.com/joserafaelSH/go-api/internal/factory"
)



func main(){
    logger.InitLogger()
    err := env.LoadEnv()
    if err == nil {
        factory.CreateFactory()
        slog.Info("Environment: " + env.ENV.GOENV)
        slog.Info("Starting WebServer on PORT: " + env.ENV.PORT)
        factory.StartGinServer()
    }
}
