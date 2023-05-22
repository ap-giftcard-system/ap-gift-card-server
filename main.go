package main

// @import
import (
	"ap-gift-card-server/db"
	"ap-gift-card-server/routers"
	"ap-gift-card-server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @notice: global variables
var (
  server			*gin.Engine
  postgresClient		*gorm.DB
)

// @dev Runs before main()
func init() {
  // load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}
  
  // set up gin engine
  server = gin.Default()

  // Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
  server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

  // init postgres client
  postgresClient = db.EstablishPostgresClient()
}

// @dev Root function
func main() {
  // Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

  // init basePath
	giftsBasePath := server.Group("/v1/ap/gifts/")

  // init Handler
  routers.ApRouter(giftsBasePath)

  // run gin server engine
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("LOCAL_DEV_PORT"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}
