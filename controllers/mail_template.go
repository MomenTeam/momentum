package controllers

import (
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/models/enums"
	"net/http"
)

type MailTemplateForm struct {
	MailTemplateType enums.MailTemplateType `json:"mailTemplateType"`
	Template string `json:"template"`
}

// CreateMailTemplate godoc
// @Summary Creates mail templates
// @Tags mailTemplate
// @Produce json
// @Param mailTemplate body MailTemplateForm true "Mail template information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /v1/mailTemplates [post]
func CreateMailTemplate(c *gin.Context) {
	mailTemplateForm := &MailTemplateForm{}
	c.BindJSON(&mailTemplateForm)

	mailTemplate := &models.MailTemplate{}

	mapper.AutoMapper(mailTemplateForm, mailTemplate)

	result, err := models.CreateMailTemplate(*mailTemplate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Mail template cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Mail template successfully created",
		"data":    result,
	})
	return
}
