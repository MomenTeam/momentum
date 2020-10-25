package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
)

// type NeedyForm struct {
// 	FirstName       string         `bson:"firstName" json:"firstName"`
// 	LastName        string         `bson:"lastName" json:"lastName"`
// 	PhoneNumber     string         `bson:"phoneNumber" json:"phoneNumber"`
// 	Summary         string         `bson:"summary" json:"summary"`
// 	Priority        int            `bson:"priority" json:"priority"`
// 	Address         models.Address `bson:"address" json:"address"`
// 	NeedyCategories []int          `bson:"needyCategories" json:"needyCategories"`
// }

// Needer type
type NeederForm struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	Category    string `json:"category"`
	PhoneNumber string `json:"phoneNumber"`
	Summary     string `json:"summary"`
}

// // GetAllNeeders godoc
// // @Summary Lists all needies informations
// // @Tags needer
// // @Produce  json
// // @Success 200 {object} gin.H
// // @Failure 400 {object} gin.H
// // @Router /v1/needer [get]
// func GetAllNeeders(c *gin.Context) {
// 	needies, _ := models.GetAllNeedies()

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"count":   len(needies),
// 		"message": "All needies informations listed",
// 		"data":    needies,
// 	})
// 	return
// }

func CreateNeeder(c *gin.Context) {
	neederForm := &NeederForm{}
	err := c.BindJSON(&neederForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer cannot be created",
			"error":   err.Error(),
		})
		return
	}

	needer := &models.Needer{
		FirstName:   neederForm.FirstName,
		LastName:    neederForm.LastName,
		PhoneNumber: neederForm.PhoneNumber,
		Summary:     neederForm.Summary,
		Address:     neederForm.Address,
		CreatedBy:   "admin", //TODO: edit this
		CreatedAt:   time.Now(),
	}

	result, err := models.CreateNeeder(*needer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Needer successfully created",
		"data":    result,
	})
	return
}