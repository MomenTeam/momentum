package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/models/enums"
	"net/http"
	"time"
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
	err := c.BindJSON(&needyForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needy cannot be created",
			"error":   err.Error(),
		})
		return
	}

	needy := &models.Needy{
		FirstName:       needyForm.FirstName,
		LastName:        needyForm.LastName,
		PhoneNumber:     needyForm.PhoneNumber,
		Summary:         needyForm.Summary,
		Priority:        needyForm.Priority,
		Address:         needyForm.Address,
		CreatedBy:       "", //TODO: edit this
		CreatedAt:       time.Now(),
	}

	for _, needyCategory := range needyForm.NeedyCategories {
		needy.NeedyCategories = append(needy.NeedyCategories, enums.GenerateNeedyCategoryFromInt(needyCategory))
	}

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

// GetAllNeediesInformations godoc
// @Summary Lists all needies informations
// @Tags needy
// @Produce  json
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needies/informations [get]
func GetAllNeediesInformations(c *gin.Context) {
	needies, _ := models.GetAllNeediesInformations()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(needies),
		"message": "All needies informations listed",
		"data":    needies,
	})
	return
}

// GetNeedyDetail godoc
// @Summary Get needy detail
// @Tags needy
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needies/getNeedyDetail/{id} [get]
func GetNeedyDetail(c *gin.Context) {
	needyId := c.Param("id")
	needyDetail, _ := models.GetNeedyDetail(needyId)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get needy detail",
		"data":    needyDetail,
	})
	return
}

// AddNeed godoc
// @Summary Add need to needy
// @Tags needy
// @Produce json
// @Param id path string true "ID"
// @Param needy body NeedForm true "Need information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/needies/{id}/addNeed [post]
func AddNeed(c *gin.Context) {
	needId := c.Param("id")
	needForm := &NeedForm{}
	err := c.BindJSON(&needForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needy cannot be created",
			"error":   err.Error(),
		})
		return
	}

	var lineItems []models.LineItem
	for _, lineItem := range needForm.LineItems {
		lineItems = append(lineItems, models.LineItem{
			Description: lineItem.Description,
			Amount:      lineItem.Amount,
			Good:        models.Good{
				Name:         lineItem.Good.Name,
				Price:        lineItem.Good.Price,
				PhotoLink:    lineItem.Good.PhotoLink,
				IsAvailable:  lineItem.Good.IsAvailable,
				GoodCategory: lineItem.Good.GoodCategory,
			},
		})
	}

	need := &models.Need{
		Name:        needForm.Name,
		Description: needForm.Description,
		LineItems:   lineItems,
		IsFulfilled: false,
		Priority:    needForm.Priority,
		IsCancelled: false,
		CreatedAt: time.Now(),
	}

	result, err := models.AddNeed(*need, needId)
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
