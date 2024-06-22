package main

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/migrate"
	"GO-JWT-Auth/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	migrate.Migrate()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080

}
