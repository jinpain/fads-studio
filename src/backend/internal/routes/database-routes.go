package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinpain/fads-studio/internal/handlers"
)

func Database(rg *gin.RouterGroup) {
	rg.GET("/getall", handlers.GetAllDatabases)
	rg.POST("/", handlers.CreateDatabase)
}
