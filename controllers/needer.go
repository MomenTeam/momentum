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

// NeederForm type
type NeederForm struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	Category    string `json:"category"`
	PhoneNumber string `json:"phoneNumber"`
	Summary     string `json:"summary"`
}

// PackageForm type
type PackageForm struct {
	NeederID string `json:"neederId"`
	Name     string `json:"name"`
}

// LineItemForm type
type LineItemForm struct {
	NeederID  string          `json:"neederId"`
	PackageID string          `json:"packageId"`
	LineItem  models.LineItem `json:"lineItem"`
}

// IsPublished type
type IsPublished struct {
	NeederID    string `json:"neederId"`
	PackageID   string `json:"packageId"`
	IsPublished bool   `json:"isPublished"`
}

type ContactForm struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PackageId   string `json:"packageId"`
	Description string `json:"description"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	NeederId    string `json:"neederId"`
}

// PackageDelete type
type PackageDelete struct {
	PackageID string `json:"packageId"`
	NeederID  string `json:"neederId"`
}

// LineItemDelete type
type LineItemDelete struct {
	PackageID  string `json:"packageId"`
	NeederID   string `json:"neederId"`
	LineItemID string `json:"lineItemId"`
}

// GetAllNeeders func
func GetAllNeeders(c *gin.Context) {
	needer, err := models.GetAllNeeders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needers couldn't fetch",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(needer),
		"message": "All needers fetched.",
		"data":    needer,
	})
	return
}

// CreateNeeder func
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
		Category:    neederForm.Category,
		Packages:    []models.Package{},
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

func NeederDetail(c *gin.Context) {
	neederID := c.Param("id")

	result, err := models.GetNeederDetail(neederID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer detail error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Needer detail.",
		"data":    result,
	})
	return
}

func CreatePackage(c *gin.Context) {
	packageForm := &PackageForm{}

	err := c.BindJSON(&packageForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer cannot be created",
			"error":   err.Error(),
		})
		return
	}

	packageModel := &models.Package{
		Name:       packageForm.Name,
		TotalPrice: 0,
	}

	result, err := models.CreatePackage(packageForm.NeederID, *packageModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Package cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Package successfully created",
		"data":    result,
	})
	return
}

func CreateLineItem(c *gin.Context) {
	lineItemForm := &LineItemForm{}

	err := c.BindJSON(&lineItemForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer cannot be created",
			"error":   err.Error(),
		})
		return
	}

	result, err := models.CreateLineItem(lineItemForm.NeederID, lineItemForm.PackageID, lineItemForm.LineItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Line Item cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Line Item successfully created",
		"data":    result,
	})
	return
}

// UpdatePublishStatusOfPackage func
func UpdatePublishStatusOfPackage(c *gin.Context) {
	isPublished := &IsPublished{}

	err := c.BindJSON(&isPublished)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "isPublished fields not correct.",
			"error":   err.Error(),
		})
		return
	}

	result, err := models.UpdatePackageIsPublished(isPublished.NeederID, isPublished.PackageID, isPublished.IsPublished)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "UpdatePublishStatusOfPackage error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Package's isPublished toggled.",
		"data":    result,
	})
	return
}

// UpdatePublishStatusOfNeeder func
func UpdatePublishStatusOfNeeder(c *gin.Context) {
	isPublished := &IsPublished{}

	err := c.BindJSON(&isPublished)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "isPublished fields not correct.",
			"error":   err.Error(),
		})
		return
	}

	result, err := models.UpdateNeederIsPublished(isPublished.NeederID, isPublished.IsPublished)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "UpdatePublishStatusOfPackage error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Needer's isPublished toggled.",
		"data":    result,
	})
	return
}

// GetAllNeedersInformation func
func GetAllNeedersInformation(c *gin.Context) {
	neederInformations, err := models.GetAllNeediesInformations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needers couldn't fetch",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(neederInformations),
		"message": "All needer information fetched.",
		"data":    neederInformations,
	})
	return
}

func GetNeederDetailAsUser(c *gin.Context) {
	neederID := c.Param("id")

	result, err := models.GetNeederDetailAsUser(neederID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Needer detail error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Needer detail.",
		"data":    result,
	})
	return
}

// CreateContact func
func CreateContact(c *gin.Context) {
	contactForm := &ContactForm{}
	err := c.BindJSON(&contactForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Contact cannot be created",
			"error":   err.Error(),
		})
		return
	}

	contact := &models.Contact{
		FirstName:   contactForm.FirstName,
		LastName:    contactForm.LastName,
		PackageId:   contactForm.PackageId,
		Description: contactForm.Description,
		PhoneNumber: contactForm.PhoneNumber,
		Email:       contactForm.Email,
		NeederId:    contactForm.NeederId,
	}

	result, err := models.CreateContact(*contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Contact cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Contact successfully created",
		"data":    result,
	})
	return
}

func GetContactRequests(c *gin.Context) {
	status := c.Param("status")

	result, err := models.GetAllContactFormsWithPackage(status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Contact requests fetch error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Contact requests",
		"data":    result,
	})
	return
}

// UpdateContactStatus func
func UpdateContactStatus(c *gin.Context) {
	contactId := c.Param("contactId")

	result, err := models.UpdateContactRequest(contactId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "UpdateContactStatus error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "UpdateContactStatus completed.",
		"data":    result,
	})
	return
}

// DeletePackage func
func DeletePackage(c *gin.Context) {
	packageDeleteRequest := &PackageDelete{}

	err := c.BindJSON(&packageDeleteRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "You sent wrong fields.",
			"error":   err.Error(),
		})
		return
	}

	result, err := models.DeletePackage(packageDeleteRequest.NeederID, packageDeleteRequest.PackageID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Package delete error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Package deleted",
		"data":    result,
	})
	return
}

// // DeleteLineItem func
// func DeleteLineItem(c *gin.Context) {
// 	lineItemDeleteRequest := &LineItemDelete{}

// 	err := c.BindJSON(&lineItemDeleteRequest)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  http.StatusBadRequest,
// 			"message": "You sent wrong fields.",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result, err := models.DeleteLineItem(lineItemDeleteRequest.NeederID, lineItemDeleteRequest.PackageID, lineItemDeleteRequest.LineItemID)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  http.StatusBadRequest,
// 			"message": "Line Item delete error",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Line Item deleted",
// 		"data":    result,
// 	})
// 	return
// }
