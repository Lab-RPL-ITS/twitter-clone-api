package repository

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	LikesRepository interface {
		LikePostById(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error
		CheckLikedPost(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error
		ErrUnlikePostById(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error
	}

	likesRepository struct {
		db *gorm.DB
	}
)

func NewLikesRepository(db *gorm.DB) LikesRepository {
	return &likesRepository{
		db: db,
	}
}

func (r *likesRepository) LikePostById(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error {
	if tx == nil {
		tx = r.db
	}

	likes := &entity.Like{
		UserID: uuid.MustParse(userId),
		PostID: postId,
	}

	if err := tx.WithContext(ctx).Create(&likes).Error; err != nil {
		return err
	}

	return nil
}

func (r *likesRepository) CheckLikedPost(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error {
	if tx == nil {
		tx = r.db
	}

	likes := &entity.Like{
		UserID: uuid.MustParse(userId),
		PostID: postId,
	}

	if err := tx.WithContext(ctx).Where("user_id = ? AND post_id = ?", userId, postId).First(&likes).Error; err != nil {
		return err
	}
	return nil
}

func (r *likesRepository) ErrUnlikePostById(ctx context.Context, tx *gorm.DB, postId uint64, userId string) error {
	if tx == nil {
		tx = r.db
	}

	likes := &entity.Like{
		UserID: uuid.MustParse(userId),
		PostID: postId,
	}

	if err := tx.WithContext(ctx).Where("user_id = ? AND post_id = ?", userId, postId).Delete(&likes).Error; err != nil {
		return err
	}

	return nil
}
