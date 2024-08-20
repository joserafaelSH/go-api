package factory

import (
	"github.com/gin-gonic/gin"
	"github.com/joserafaelSH/go-api/config/env"
	"github.com/joserafaelSH/go-api/internal/handler"
	healthhandler "github.com/joserafaelSH/go-api/internal/handler/health_handler"
	"github.com/joserafaelSH/go-api/internal/service"
	healthservice "github.com/joserafaelSH/go-api/internal/service/health_service"
)


type Factory struct {
    HealthService service.ServiceInterface[any, string]
    HealthHandler handler.HandlerInterface
}


var FACTORY Factory

func CreateFactory(){
    healthService := healthservice.CreateHealthService()
    healthHandler := healthhandler.CreateHealthHandler(healthService)
    FACTORY = Factory{
        HealthService: healthService,
        HealthHandler: healthHandler,
    }
}

func StartGinServer(){
    r := gin.Default()

    r.GET("/health", func(ctx *gin.Context) {
        FACTORY.HealthHandler.Execute(ctx)
    })
    port := env.ENV.PORT
    err := r.Run(":" + port)
    if err != nil {
        panic("[Error] failed to start Gin server due to: " + err.Error())
    }
}
