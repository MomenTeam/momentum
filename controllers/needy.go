package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
	"net/http"
)

type NeedyForm struct {
	FirstName       string `json:"firstName"`
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

	fmt.Println(needyForm)

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