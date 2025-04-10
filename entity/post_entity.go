package entity

import "github.com/google/uuid"

type Post struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Text string    `gorm:"not null" json:"text"`

	Parent   *Post      `gorm:"foreignkey:ParentID" json:"parent,omitempty"`
	ParentID *uuid.UUID `gorm:"type:uuid" json:"parent_id,omitempty"`

	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	User   User      `gorm:"foreignkey:UserID" json:"user"`

	Timestamp
}
