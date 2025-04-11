package seeds

import (
	"encoding/json"
	"errors"
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	var users []entity.User
	err = db.Find(&users).Limit(2).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	postMap := make(map[int]*entity.Post)

	for i, post := range listPost {
		newPost := entity.Post{
			Text:   post.Text,
			UserID: users[i%len(users)].ID,
		}

		if err := db.Create(&newPost).Error; err != nil {
			return err
		}
		postMap[i] = &newPost
	}

	for i, post := range listPost {
		if post.ParentID != nil {
			parentPost := postMap[int(*post.ParentID-1)]
			if parentPost != nil {
				currentPost := postMap[i]
				currentPost.ParentID = &parentPost.ID
				if err := db.Save(currentPost).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
