package provider

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/controller"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvidePostDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)

	// Service
	postService := service.NewPostService(userRepository, postRepository, jwtService)

	// Controller
	do.Provide(injector, func(i *do.Injector) (controller.PostController, error) {
		return controller.NewPostController(postService), nil
	})
}
