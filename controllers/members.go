package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/prashantvnsi/goAPI/models"
	//"../models"
)

// GET /members
// Get all members
func FindMembers(c *gin.Context) {
	var members []models.Members
	models.DB.Find(&members)
  
	c.JSON(http.StatusOK, gin.H{"data": members})
  }