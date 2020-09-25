package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssigneeForm struct {
	Name       string `json:"name"`
	Department string `json:"department" validate:"department"`
}

// CreateAssignee godoc
// @Summary Creates assignee
// @Tags assignee
// @Produce  json
// @Param assignee body AssigneeForm true "Assignee information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /assignees [post]
func CreateAssignee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Assignee successfully created",
		"data":    "asd", //TODO: change this
	})
	return
}
