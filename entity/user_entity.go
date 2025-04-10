package entity

import (
	"github.com/Lab-RPL-ITS/twitter-clone-api/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Username string    `gorm:"not null" gorm:"unique" json:"username"`
	Bio      string    `gorm:"not null" json:"bio"`
	Password string    `gorm:"not null" json:"password"`
	ImageUrl *string   `json:"image_url"`

	Posts []Post `gorm:"foreignkey:UserID" json:"posts,omitempty"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
