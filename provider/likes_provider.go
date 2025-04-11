package provider

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/controller"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideLikesDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	likesRepository := repository.NewLikesRepository(db)
	postRepository := repository.NewPostRepository(db)

	// Service
	likesService := service.NewLikesService(likesRepository, postRepository, jwtService)

	// Controller
	do.Provide(injector, func(i *do.Injector) (controller.LikesController, error) {
		return controller.NewLikesController(likesService), nil
	})
}
