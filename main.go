package main

import (
	r "GOLANG/api/router"
	"GOLANG/config"
	docs "GOLANG/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDb()
	config.SyncDatabase()
}

// @title           Simple JWT Auth API With Gin Framework
// @version         1.0
// @contact.name   Yasin Çakır
// @contact.url    https://www.linkedin.com/in/yasincakir26/
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey JwtAuth(http, Bearer)
// @in header
// @name JWT Authorization
func main() {

	// Gin Framework
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	r.GetRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")

}
