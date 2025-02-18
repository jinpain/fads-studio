package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinpain/fads-studio/internal/configuration"

	docs "github.com/jinpain/fads-studio/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FADS Studio WebAPI
// @version 1.0
// @description ...

// @host localhost:8080
// @BasePath /api/v1
func main() {
	config, err := configuration.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		v1.GET("")
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(config.Server.Host + ":" + config.Server.Port); err != nil {
		panic(err)
	}
}
