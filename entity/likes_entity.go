package entity

import "github.com/google/uuid"

type Like struct {
	PostID uint64 `gorm:"primaryKey;not null" json:"post_id"`
	Post   Post   `gorm:"foreignkey:PostID" json:"post"`

	UserID uuid.UUID `gorm:"primaryKey;not null" json:"user_id"`
	User   User      `gorm:"foreignkey:UserID" json:"user"`

	Timestamp
}
