package router

import (
	"GOLANG/api/controllers"
	"GOLANG/api/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {

	// User Route
	api_users := r.Group("/api/users")
	api_users.POST("/register", controllers.Register)
	api_users.POST("/auth", controllers.Auth)
	api_users.GET("/me", middlewares.CheckAuth, controllers.GetUserProfile)

}
