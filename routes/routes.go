package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/momenteam/momentum/controllers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", welcome)
	router.NoRoute(notFound)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 24 * time.Hour,
	}))

	needies := router.Group("/v1/needies")
	{
		needies.POST("/", controllers.CreateNeedy)
		needies.GET("/", controllers.GetAllNeedies)
		needies.GET("/informations", controllers.GetAllNeediesInformations)
	}

	mailTemplates := router.Group("/v1/mailTemplates")
	{
		mailTemplates.POST("/", controllers.CreateMailTemplate)
	}

	/*
		assignees := router.Group("/assignees")
		{
			assignees.POST("/", controllers.CreateAssignee)
			assignees.GET("/", controllers.GetAllNeedies)
			assignees.GET("/find-by-id/:id", controllers.GetAssignee) // This is not a good practice but I have to
			assignees.GET("/find-by-name/:name", controllers.GetAssigneeByName)
			assignees.DELETE("/:id", controllers.DeleteAssignee)
		}
	*/
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Momentum API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
