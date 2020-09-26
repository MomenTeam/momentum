package main

import (
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/configs"
	"github.com/momenteam/momentum/controllers"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/docs"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/routes"
	"log"
)

func init() {
	configs.Setup()
	database.Setup()
	mapper.Register(&models.MailTemplate{})
	mapper.Register(&controllers.MailTemplateForm{})
}

func main() {
	docs.SwaggerInfo.Title = "Momentum"
	docs.SwaggerInfo.Description = "Momentum"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + configs.GlobalConfig.Server.Port
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":" + configs.GlobalConfig.Server.Port))
}