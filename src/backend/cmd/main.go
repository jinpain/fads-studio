package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinpain/fads-studio/internal/configuration"
	"github.com/jinpain/fads-studio/internal/routes"

	docs "github.com/jinpain/fads-studio/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//var config *configuration.Config

// @title FADS Studio WebAPI
// @version 1.0
// @description ...

// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := configuration.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(configuration.Config)

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		routes.Database(v1.Group("/database"))
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(configuration.Config.Server.Host + ":" + configuration.Config.Server.Port); err != nil {
		panic(err)
	}
}
