package main

import (
	"GOLANG/api/initializers"
	r "GOLANG/api/router"
	docs "GOLANG/docs"
	"os"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// @title           Simple JWT Auth API With Gin Framework
// @version         1.0
// @contact.name   Yasin Çakır
// @contact.url    https://www.linkedin.com/in/yasincakir26/
// @host      127.0.0.1:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// Gin Framework
	router := gin.Default()

	// Cors Settings
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsConfig)
	docs.SwaggerInfo.BasePath = "/api"

	r.GetRoute(router)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + os.Getenv("PORT"))

}
