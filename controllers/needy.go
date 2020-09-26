package controllers

import (
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/models/enums"
	"net/http"
)

type NeedyForm struct {
	FirstName     string                  `bson:"firstName" json:"firstName"`
	LastName      string                  `bson:"lastName" json:"lastName"`
	PhoneNumber   string                  `bson:"phoneNumber" json:"phoneNumber"`
	Summary       string                  `bson:"summary" json:"summary"`
	Priority      int                     `bson:"priority" json:"priority"`
	Address       models.Address          `bson:"address" json:"address"`
	NeedyCategory enums.NeedyCategoryType `bson:"needyCategory" json:"needyCategory"`
	Needs         []models.Need           `bson:"needs" json:"needs"`
}

// CreateNeedy godoc
// @Summary Creates needy
// @Tags needy
// @Produce json
// @Param needy body NeedyForm true "Needy information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needies [post]
func CreateNeedy(c *gin.Context) {
	needyForm := &NeedyForm{}
	c.BindJSON(&needyForm)

	needy := &models.Needy{}

	mapper.AutoMapper(needyForm, needy)

	result, err := models.CreateNeedy(*needy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needy cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Needy successfully created",
		"data":    result,
	})
	return
}

// GetAllNeedies godoc
// @Summary Lists all needies
// @Tags needy
// @Produce  json
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needies [get]
func GetAllNeedies(c *gin.Context) {
	needies, _ := models.GetAllNeedies()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(needies),
		"message": "All needies listed",
		"data":    needies,
	})
	return
}
