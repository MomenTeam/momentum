package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devfeel/mapper"
	"github.com/momenteam/momentum/models"
	"github.com/momenteam/momentum/validator"
	"github.com/gin-gonic/gin"
)

type CandidateForm struct {
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email" validate:"email"`
	Department      string    `json:"department" validate:"department"`
	University      string    `json:"university"`
	Experience      bool      `json:"experience"`
	Assignee        string    `json:"assignee"`
	ApplicationDate time.Time `json:"application_date"`
}

type ArrangeMeetingForm struct {
	NextMeetingTime *time.Time `json:"next_meeting_time"`
}

// CreateCandidate godoc
// @Summary Creates candidate
// @Tags candidate
// @Produce  json
// @Param candidate body CandidateForm true "Candidate information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /candidates [post]
func CreateCandidate(c *gin.Context) {
	candidateForm := &CandidateForm{}
	c.BindJSON(&candidateForm)

	candidate := &models.Candidate{}

	fmt.Println(candidateForm)

	mapper.AutoMapper(candidateForm, candidate)

	validationError := validator.Validator.Struct(*candidateForm)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Validation error",
			"error":   validationError.Error(),
		})
		return
	}

	result, err := models.CreateCandidate(*candidate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Candidate cannot be created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Candidate successfully created",
		"data":    result,
	})
	return
}

// GetAllCandidates godoc
// @Summary Lists all candidates
// @Tags candidate
// @Produce  json
// @Success 200 {array} models.Candidate
// @Failure 400 {object} gin.H
// @Router /candidates [get]
func GetAllCandidates(c *gin.Context) {
	candidates, _ := models.GetAllCandidates()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"count":   len(candidates),
		"message": "All Candidates listed",
		"data":    candidates,
	})
	return
}

// GetCandidate godoc
// @Summary Gets candidate by id
// @Tags candidate
// @Produce  json
// @Success 200 {array} models.Candidate
// @Param id path string true "ID"
// @Failure 400 {object} gin.H
// @Router /candidates/{id} [get]
func GetCandidate(c *gin.Context) {
	id := c.Param("id")

	candidate, err := models.ReadCandidate(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Candidate cannot be found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Candidate found",
		"data":    candidate,
	})
	return
}

// DeleteCandidate godoc
// @Summary Deletes candidate
// @Tags candidate
// @Produce  json
// @Param id path string true "ID"
// @Failure 400 {object} gin.H
// @Router /candidates/{id} [delete]
func DeleteCandidate(c *gin.Context) {
	id := c.Param("id")

	err := models.DeleteCandidate(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Candidate cannot be deleted",
			"error":   err.Error(),
		})
		return
	}
}

// ArrangeMeeting godoc
// @Summary Arrangenges meeting
// @Tags candidate
// @Produce  json
// @Param id path string true "ID"
// @Param candidate body ArrangeMeetingForm true "Candidate information"
// @Failure 400 {object} gin.H
// @Router /candidates/{id}/arrange-meeting [put]
func ArrangeMeeting(c *gin.Context) {
	id := c.Param("id")

	form := &ArrangeMeetingForm{}
	c.BindJSON(&form)

	err := models.ArrangeMeeting(id, form.NextMeetingTime)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Meeting cannot be arrenged",
			"error":   err.Error(),
		})
		return
	}
}

// CompleteMeeting godoc
// @Summary Completes meeting
// @Tags candidate
// @Produce  json
// @Param id path string true "ID"
// @Failure 400 {object} gin.H
// @Router /candidates/{id}/complete-meeting [put]
func CompleteMeeting(c *gin.Context) {
	id := c.Param("id")

	err := models.CompleteMeeting(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Meeting cannot be completed",
			"error":   err.Error(),
		})
		return
	}
}

// AcceptCandidate godoc
// @Summary Accept
// @Tags candidate
// @Produce  json
// @Param id path string true "ID"
// @Failure 400 {object} gin.H
// @Router /candidates/{id}/accept [put]
func AcceptCandidate(c *gin.Context) {
	id := c.Param("id")

	err := models.AcceptCandidate(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Candidate cannot be accepted",
			"error":   err.Error(),
		})
		return
	}
}

// DenyCandidate godoc
// @Summary Denies meeting
// @Tags candidate
// @Produce  json
// @Param id path string true "ID"
// @Failure 400 {object} gin.H
// @Router /candidates/{id}/deny [put]
func DenyCandidate(c *gin.Context) {
	id := c.Param("id")

	err := models.DenyCandidate(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Candidate cannot be accepted",
			"error":   err.Error(),
		})
		return
	}
}
