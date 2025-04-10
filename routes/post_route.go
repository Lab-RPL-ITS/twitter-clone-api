package routes

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/controller"
	"github.com/Lab-RPL-ITS/twitter-clone-api/middleware"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Post(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	postController := do.MustInvoke[controller.PostController](injector)

	routes := route.Group("/api/post")
	{
		routes.POST("/", middleware.Authenticate(jwtService), postController.Create)
	}
}
