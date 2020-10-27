package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", welcome)
	router.NoRoute(notFound)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	needies := router.Group("/v1/needer")
	{
		needies.GET("/all", controllers.GetAllNeeders)
		needies.POST("/create", controllers.CreateNeeder)
		needies.GET("/detail/:id", controllers.NeederDetail)
		needies.POST("/createPackage", controllers.CreatePackage)
		needies.POST("/createLineItem", controllers.CreateLineItem)
		needies.PUT("/updatePublishStatusOfPackage", controllers.UpdatePublishStatusOfPackage)
		needies.PUT("/updatePublishStatus", controllers.UpdatePublishStatusOfNeeder)
		needies.PUT("/getAllNeederInformations", controllers.GetAllNeedersInformation)
	}

	// needs := router.Group("/v1/needs")
	// {
	// 	needs.GET("/", controllers.GetAllNeeds)
	// 	needs.GET("/getAllDetails", controllers.GetAllNeedDetails)
	// 	needs.GET("/setFulfilled/:needId", controllers.SetFulfilled)
	// 	needs.DELETE("/:needId/cancel", controllers.CancelNeed)
	// }

	// payment := router.Group("/v1/payment")
	// {
	// 	payment.POST("/:needId", controllers.PayForNeed)
	// }

	// mailTemplates := router.Group("/v1/mailTemplates")
	// {
	// 	mailTemplates.POST("/", controllers.CreateMailTemplate)
	// }

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
