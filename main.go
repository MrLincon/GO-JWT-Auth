package main

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/migrate"
	"GO-JWT-Auth/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	migrate.Migrate()
}

func main() {

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Info("Server is running on port 3000")

	r := gin.Default()
	routes.SetupRoutes(r, sugar)

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		logger.Error("PORT is not set")
		return
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
	// Listen and serve on 0.0.0.0:3000

}
