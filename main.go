package main

// @import
import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

// @dev Root function
func main() {
  // Loads environment variables
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  // Init gin engine
  r := gin.Default()

  // HTTP Get
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // run gin engine
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
