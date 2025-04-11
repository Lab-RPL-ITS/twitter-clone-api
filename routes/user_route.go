package routes

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/controller"
	"github.com/Lab-RPL-ITS/twitter-clone-api/middleware"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func User(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	userController := do.MustInvoke[controller.UserController](injector)

	routes := route.Group("/api/user")
	{
		// User
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		routes.POST("/check-username", userController.CheckUsername)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		routes.GET("/:username", userController.GetUserByUsername)
		routes.GET("/:username/posts", userController.GetUserPosts)
		routes.PATCH("/update", middleware.Authenticate(jwtService), userController.UpdateUser)
	}
}
