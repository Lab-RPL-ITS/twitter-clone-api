package provider

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/config"
	"github.com/Lab-RPL-ITS/twitter-clone-api/constants"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	ProvideUserDependencies(injector)
	ProvidePostDependencies(injector)
	ProvideLikesDependencies(injector)
}
