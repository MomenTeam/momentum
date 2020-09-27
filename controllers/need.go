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
	CreditCardNumber string `json:"creditCardNumer"`
	Cvv string `json:"cvv"`
	ExpireDate string `json:"expireDate"`
	FullName string `json:"fullName"`
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