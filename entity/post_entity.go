package entity

import "github.com/google/uuid"

type Post struct {
	ID   uint64 `gorm:"primary_key;autoIncrement" json:"id"`
	Text string `gorm:"not null" json:"text"`

	Parent   *Post   `gorm:"foreignkey:ParentID" json:"parent,omitempty"`
	ParentID *uint64 `json:"parent_id,omitempty"`

	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	User   User      `gorm:"foreignkey:UserID" json:"user"`

	Timestamp
}
