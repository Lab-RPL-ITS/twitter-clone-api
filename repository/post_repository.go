package repository

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	PostRepository interface {
		CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
		GetPostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID) (entity.Post, error)
	}

	postRepository struct {
		db *gorm.DB
	}
)

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&post).Error; err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (r *postRepository) GetPostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	var post entity.Post
	if err := tx.WithContext(ctx).Joins("JOIN users ON users.id = posts.user_id ").Where("posts.id = ?", postId).Take(&post).Error; err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
