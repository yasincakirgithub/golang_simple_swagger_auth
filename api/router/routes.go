package router

import (
	"GOLANG/api/controllers"
	"GOLANG/api/middleware"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {

	// User Route
	api_users := r.Group("/api/users")
	api_users.POST("/signup", controllers.Signup)
	api_users.POST("/login", controllers.Login)
	api_users.POST("/logout", middleware.RequireAuth, controllers.Logout)

}
