package controllers

import (
	"github.com/gin-gonic/gin"
	"models"
)

// GET /group
// Get all groups
func FindGroups(c *gin.Context) {
	var groups []models.Group
	models.DB.Find(&groups)
  
	c.JSON(http.StatusOK, gin.H{"data": groups})
  }