package migrations

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Like{},
	); err != nil {
		return err
	}

	return nil
}
