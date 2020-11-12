package main

import (
  "github.com/gin-gonic/gin"
  "github.com/rahmanfadhil/gin-bookstore/models"
  "github.com/rahmanfadhil/gin-bookstore/controllers"
)

func main() {
  r := gin.Default()

  /*r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })*/

  models.ConnectDatabase()

  r.GET("/groups", controllers.FindGroups)

  r.Run()
}