package controllers

import (
	"event-app/config"
	"event-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event after being authenticated
// @Tags events
// @Accept  json
// @Produce  json
// @Param   event body models.Event true "Event"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /events [post]
// @Security Bearer
func CreateEvent(c *gin.Context) {
	username := c.MustGet("username").(string)
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.UserID = user.ID
	if err := config.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event created"})
}

// GetEvents godoc
// @Summary Get all events
// @Description Retrieve all events in the database
// @Tags events
// @Produce  json
// @Success 200 {array} models.Event
// @Router /events [get]
// @Security Bearer
func GetEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

// GetTodayEvents godoc
// @Summary Get today's events
// @Description Retrieve events occurring today
// @Tags events
// @Produce  json
// @Success 200 {array} models.Event
// @Router /events/today [get]
// @Security Bearer
func GetTodayEvents(c *gin.Context) {
	var events []models.Event
	today := time.Now().Format("2006-01-02")
	config.DB.Where("event_date = ?", today).Find(&events)
	c.JSON(http.StatusOK, events)
}
