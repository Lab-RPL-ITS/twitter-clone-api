package repository

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	PostRepository interface {
		CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
		GetPostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID) (entity.Post, error)
		DeletePostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID) error
		UpdatePostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID, post entity.Post) (entity.Post, error)
		GetAllPostsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPostsRepositoryResponse, error)
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
	if err := tx.WithContext(ctx).Joins("User").Where("posts.id = ?", postId).Take(&post).Error; err != nil {
		return entity.Post{
			ID:       post.ID,
			Text:     post.Text,
			ParentID: post.ParentID,
			User: entity.User{
				ID:       post.User.ID,
				Name:     post.User.Name,
				Username: post.User.Username,
				Bio:      post.User.Bio,
				ImageUrl: post.User.ImageUrl,
			},
		}, err
	}

	return post, nil
}

func (r *postRepository) DeletePostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Post{}, postId).Error; err != nil {
		return err
	}

	return nil
}

func (r *postRepository) UpdatePostById(ctx context.Context, tx *gorm.DB, postId uuid.UUID, post entity.Post) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&entity.Post{}).Where("id = ?", postId).Updates(post).Error; err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (r *postRepository) GetAllPostsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPostsRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var posts []entity.Post
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Post{}).Joins("User")
	if req.Search != "" {
		query = query.Where("text LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllPostsRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&posts).Error; err != nil {
		return dto.GetAllPostsRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllPostsRepositoryResponse{
		Posts: posts,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
