package routes

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/controller"
	"github.com/Lab-RPL-ITS/twitter-clone-api/middleware"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Likes(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	likesController := do.MustInvoke[controller.LikesController](injector)

	routes := route.Group("/api/likes")
	{
		routes.PUT("/:post_id", middleware.Authenticate(jwtService), likesController.LikePostById)
		routes.DELETE("/:post_id", middleware.Authenticate(jwtService), likesController.UnlikePostById)
	}
}
