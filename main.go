package main

import (
  "github.com/gin-gonic/gin"
  //"./controllers"
  //"./models"
  "github.com/prashantvnsi/goAPI/models"
  "github.com/prashantvnsi/goAPI/controllers"
)

func main() {
  r := gin.Default()

  /*r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })*/
  models.ConnectDatabase()
  //models.ConnectDatabase()

  r.GET("/groups", controllers.FindGroups)

  r.Run()
}