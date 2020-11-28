package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/prashantvnsi/goAPI/models"
	//"../models"
)

// GET /group
// Get all groups
func FindGroups(c *gin.Context) {
	var groups []models.Group
	models.DB.Find(&groups)
  
	c.JSON(http.StatusOK, gin.H{"data": groups})
  }