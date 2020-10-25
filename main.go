package main

import (
	"log"

	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/configs"
	"github.com/momenteam/momentum/controllers"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/docs"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/routes"
)

func init() {
	configs.Setup()
	database.Setup()
	mapper.Register(&models.MailTemplate{})
	mapper.Register(&controllers.NeederForm{})
}

func main() {
	docs.SwaggerInfo.Title = "Momentum"
	docs.SwaggerInfo.Description = "Momentum"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = configs.GlobalConfig.AppBaseBath
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":" + configs.GlobalConfig.Server.Port))
}
