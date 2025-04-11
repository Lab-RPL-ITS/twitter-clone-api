package repository

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"gorm.io/gorm"
)

type (
	PostRepository interface {
		CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
		GetPostById(ctx context.Context, tx *gorm.DB, postId uint64) (entity.Post, error)
		DeletePostById(ctx context.Context, tx *gorm.DB, postId uint64) error
		UpdatePostById(ctx context.Context, tx *gorm.DB, postId uint64, post entity.Post) (entity.Post, error)
		GetAllPostsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPostsRepositoryResponse, error)
		GetAllPostsWithPaginationByUsername(ctx context.Context, tx *gorm.DB, username string, req dto.UserPostsPaginationRequest) (dto.GetAllPostsRepositoryResponse, error)
		GetAllPostRepliesWithPagination(ctx context.Context, tx *gorm.DB, postId uint64, req dto.PaginationRequest) (dto.GetAllRepliesRepositoryResponse, error)
		UpdateLikesCount(ctx context.Context, tx *gorm.DB, postId uint64, count int) error
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

func (r *postRepository) GetPostById(ctx context.Context, tx *gorm.DB, postId uint64) (entity.Post, error) {
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

func (r *postRepository) DeletePostById(ctx context.Context, tx *gorm.DB, postId uint64) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Post{}, postId).Error; err != nil {
		return err
	}

	return nil
}

func (r *postRepository) UpdatePostById(ctx context.Context, tx *gorm.DB, postId uint64, post entity.Post) (entity.Post, error) {
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

	query := tx.WithContext(ctx).Model(&entity.Post{}).Joins("User").Unscoped().Where("posts.parent_id IS NULL").Order("created_at DESC")
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

func (r *postRepository) GetAllPostRepliesWithPagination(ctx context.Context, tx *gorm.DB, postId uint64, req dto.PaginationRequest) (dto.GetAllRepliesRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var replies []entity.Post
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Post{}).Joins("User").Unscoped().Where("posts.parent_id = ?", postId).Order("created_at DESC")
	if req.Search != "" {
		query = query.Where("text LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllRepliesRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Unscoped().Find(&replies).Error; err != nil {
		return dto.GetAllRepliesRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllRepliesRepositoryResponse{
		Replies: replies,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *postRepository) UpdateLikesCount(ctx context.Context, tx *gorm.DB, postId uint64, count int) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&entity.Post{}).Where("id = ?", postId).UpdateColumn("total_likes", gorm.Expr("total_likes + ?", count)).Error; err != nil {
		return err
	}

	return nil
}

func (r *postRepository) GetAllPostsWithPaginationByUsername(ctx context.Context, tx *gorm.DB, username string, req dto.UserPostsPaginationRequest) (dto.GetAllPostsRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var posts []entity.Post
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Post{}).Joins("User").Unscoped().Where("posts.parent_id IS NULL").Where("\"User\".username = ?", username)
	if req.IsLiked {
		query = query.Joins("INNER JOIN likes ON likes.post_id = posts.id AND likes.user_id = posts.user_id")
	}
	query = query.Order("created_at DESC")

	if req.Search != "" {
		query = query.Where("text LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllPostsRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(dto.PaginationRequest{
		Page:    req.Page,
		PerPage: req.PerPage,
	})).Unscoped().Find(&posts).Error; err != nil {
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
