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

// func GetNeedyDetail(id string) (NeedyDetail, error) {
// 	needy := Needy{}
// 	err := database.NeediesCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needy)

// 	needFilter := bson.M{"_id": bson.M{"$in": needy.Needs}}

// 	var needs []Need
// 	cursor, err := database.NeedCollection.Find(context.Background(), needFilter)

// 	if cursor != nil {
// 		for cursor.Next(context.Background()) {
// 			var need Need
// 			if err = cursor.Decode(&need); err != nil {
// 				log.Fatal(err)
// 			}
// 			needs = append(needs, need)
// 		}

// 		needyDetail := NeedyDetail{
// 			ID:              needy.ID,
// 			FirstName:       needy.FirstName,
// 			LastName:        needy.LastName,
// 			ShortName:  fmt.Sprintf("%c%c", needy.FirstName[0], needy.LastName[0]),
// 			MaskedName:        fmt.Sprintf("%s %s", mask(needy.FirstName), mask(needy.LastName)),
// 			PhoneNumber:     needy.PhoneNumber,
// 			Summary:         needy.Summary,
// 			Priority:        needy.Priority,
// 			Address:         needy.Address,
// 			NeedyCategories: needy.NeedyCategories,
// 			Needs:           needs,
// 			CreatedBy:       needy.CreatedBy,
// 			CreatedAt:       needy.CreatedAt,
// 		}

// 		return needyDetail, err
// 	}

// 	return NeedyDetail{}, err
// }