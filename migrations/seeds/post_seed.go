package seeds

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"gorm.io/gorm"
)

func ListPostSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/posts.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listPost []entity.Post
	if err := json.Unmarshal(jsonData, &listPost); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Post{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Post{}); err != nil {
			return err
		}
	}

	var user entity.User
	err = db.Find(&user).Limit(2).Error
	if err != nil {
		return err
	}

	return nil
}
