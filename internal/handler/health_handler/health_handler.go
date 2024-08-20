package healthhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/joserafaelSH/go-api/internal/service"
)

type HealthHandler struct{
    HealthService service.ServiceInterface[any, string]
}

func CreateHealthHandler(service service.ServiceInterface[any, string]) *HealthHandler{
    return &HealthHandler{
        HealthService: service,
    }

}

func (h HealthHandler) Execute(c *gin.Context){
    message:= h.HealthService.Execute("")
    c.JSON(200, gin.H{
		"message": message,

	})
}
