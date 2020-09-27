package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
	"net/http"
)

type NeedForm struct {
	Name        string         `bson:"name" json:"name"`
	Description string         `bson:"description" json:"description"`
	LineItems   []LineItemForm `bson:"lineItems" json:"lineItems"`
	Priority    int            `bson:"priority" json:"priority"`
}

type LineItemForm struct {
	Description string      `bson:"description" json:"description"`
	Amount      int         `bson:"amount" json:"amount"`
	Good        models.Good `bson:"good" json:"good"`
	//GoodIds     []int  `bson:"goodIds" json:"goodIds"`
}

type PaymentForm struct {
	CreditCardNumber string `json:"creditCardNumber"`
	Cvv              string `json:"cvv"`
	ExpireDate       string `json:"expireDate"`
	FullName         string `json:"fullName"`
}

// GetAllNeeds godoc
// @Summary Lists all needs
// @Tags need
// @Produce  json
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needs [get]
func GetAllNeeds(c *gin.Context) {
	needs, _ := models.GetAllNeeds()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(needs),
		"message": "All needs listed",
		"data":    needs,
	})
	return
}

// PayForNeed godoc
// @Summary Pay need
// @Tags need
// @Produce json
// @Param needId path string true "ID"
// @Param payment body PaymentForm true "Payment information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/payment/{needId} [post]
func PayForNeed(c *gin.Context) {
	needId := c.Param("needId")

	paymentForm := &PaymentForm{}
	err := c.BindJSON(&paymentForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Payment cannot be read",
			"error":   err.Error(),
		})
		return
	}

	payment, err := models.PayNeed(needId, paymentForm.FullName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Payment error",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Payment successful",
		"data":    payment,
	})
	return
}

// SetFulfilled godoc
// @Summary Set need as fulfilled
// @Tags need
// @Produce json
// @Param needId path string true "ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needs/{needId}/setFulfilled [get]
func SetFulfilled(c *gin.Context) {
	needId := c.Param("needId")

	_, err := models.SetFulfilled(needId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Set fulfilled error",
			"error":   err.Error(),
		})
	}

	return
}

// Cancel godoc
// @Summary Cancels need
// @Tags need
// @Produce json
// @Param needId path string true "ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needs/{needId}/cancel [delete]
func CancelNeed(c *gin.Context) {
	needId := c.Param("needId")

	_, err := models.CancelNeed(needId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Cancel error",
			"error":   err.Error(),
		})
	}

	return
}

