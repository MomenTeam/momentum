package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
)

type NeedyForm struct {
	FirstName       string         `bson:"firstName" json:"firstName"`
	LastName        string         `bson:"lastName" json:"lastName"`
	PhoneNumber     string         `bson:"phoneNumber" json:"phoneNumber"`
	Summary         string         `bson:"summary" json:"summary"`
	Priority        int            `bson:"priority" json:"priority"`
	Address         models.Address `bson:"address" json:"address"`
	NeedyCategories []int          `bson:"needyCategories" json:"needyCategories"`
}

// GetAllNeeders godoc
// @Summary Lists all needies informations
// @Tags needer
// @Produce  json
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needer [get]
func GetAllNeeders(c *gin.Context) {
	needies, _ := models.GetAllNeedies()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(needies),
		"message": "All needies informations listed",
		"data":    needies,
	})
	return
}
