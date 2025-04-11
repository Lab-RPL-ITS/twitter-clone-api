package entity

import "github.com/google/uuid"

type Post struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Text       string `gorm:"not null" json:"text"`
	TotalLikes uint64 `gorm:"default:0" json:"total_likes"`

	Parent   *Post   `gorm:"foreignkey:ParentID" json:"parent,omitempty"`
	ParentID *uint64 `json:"parent_id,omitempty"`

	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	User   User      `gorm:"foreignkey:UserID" json:"user"`

	Timestamp
}
