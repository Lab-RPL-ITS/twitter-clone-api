package migrations

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListPostSeeder(db); err != nil {
		return err
	}

	return nil
}
