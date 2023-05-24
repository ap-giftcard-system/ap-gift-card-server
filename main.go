package main

// @import
import (
	"ap-gift-card-server/controllers"
	"ap-gift-card-server/dao"
	"ap-gift-card-server/db"
	"ap-gift-card-server/routers"
	"ap-gift-card-server/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice: global variables
var (
	server			*gin.Engine
	ctx			context.Context
	mongoClient		*mongo.Client
	agr 		*routers.ApGiftRouter
)

// @dev Runs before main()
func init() {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}
	
	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init context
	ctx = context.TODO()

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// get ap-gift-holders collection
	apGiftCollection := db.GetMongoCollection(mongoClient, "ap-gift-holders")

	// init apGiftDao
	agd := dao.ApGiftDaoConstructor(ctx, apGiftCollection)

	// init apGiftController
	agc := controllers.ApGiftControllerContructor(&agd)

	// init apGiftRouter
	agr = routers.ApGiftRouterConstructor(agc)
}

// @dev Root function
func main() {
	// setup CORS
	server.Use(utils.SetupCorsConfig())

  	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()

  	// init basePath
	apGiftBasePath := server.Group("/v1/ap/gift/holder")

	// init Handler
	agr.ApRouter(apGiftBasePath)

	// run gin server engine
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("LOCAL_DEV_PORT"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}
