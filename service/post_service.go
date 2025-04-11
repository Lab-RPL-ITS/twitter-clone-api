package service

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
	"github.com/google/uuid"
)

type (
	PostService interface {
		CreatePost(ctx context.Context, userId string, req dto.PostCreateRequest) (dto.PostResponse, error)
		GetPostById(ctx context.Context, postId uint64) (dto.PostRepliesPaginationResponse, error)
		DeletePostById(ctx context.Context, postId uint64) error
		UpdatePostById(ctx context.Context, userId string, postId uint64, req dto.PostUpdateRequest) (dto.PostResponse, error)
		GetAllPosts(ctx context.Context, req dto.PaginationRequest) (dto.PostPaginationResponse, error)
	}

	postService struct {
		userRepo   repository.UserRepository
		postRepo   repository.PostRepository
		jwtService JWTService
	}
)

func NewPostService(userRepo repository.UserRepository, postRepo repository.PostRepository, jwtService JWTService) PostService {
	return &postService{
		userRepo:   userRepo,
		postRepo:   postRepo,
		jwtService: jwtService,
	}
}

func (s *postService) CreatePost(ctx context.Context, userId string, req dto.PostCreateRequest) (dto.PostResponse, error) {
	if req.ParentID != nil {
		_, err := s.postRepo.GetPostById(ctx, nil, *req.ParentID)
		if err != nil {
			return dto.PostResponse{}, dto.ErrGetPostById
		}
	}

	post := entity.Post{
		Text:     req.Text,
		UserID:   uuid.MustParse(userId),
		ParentID: req.ParentID,
	}

	result, err := s.postRepo.CreatePost(ctx, nil, post)
	if err != nil {
		return dto.PostResponse{}, dto.ErrCreatePost
	}

	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.PostResponse{}, dto.ErrGetUserById
	}

	return dto.PostResponse{
		ID:       result.ID,
		Text:     result.Text,
		ParentID: req.ParentID,
		User: dto.UserResponse{
			ID:       user.ID.String(),
			Name:     user.Name,
			Bio:      user.Bio,
			UserName: user.Username,
			ImageUrl: user.ImageUrl,
		},
	}, nil
}

func (s *postService) GetPostById(ctx context.Context, postId uint64) (dto.PostRepliesPaginationResponse, error) {
	post, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.PostRepliesPaginationResponse{}, dto.ErrGetPostById
	}

	replies, err := s.postRepo.GetAllPostRepliesWithPagination(ctx, nil, postId, dto.PaginationRequest{})
	if err != nil {
		return dto.PostRepliesPaginationResponse{}, dto.ErrGetPostReplies
	}

	var data []dto.PostResponse
	for _, reply := range replies.Replies {
		datum := dto.PostResponse{
			ID:       reply.ID,
			Text:     reply.Text,
			ParentID: reply.ParentID,
			User: dto.UserResponse{
				ID:       reply.UserID.String(),
				Name:     reply.User.Name,
				Bio:      reply.User.Bio,
				UserName: reply.User.Username,
				ImageUrl: reply.User.ImageUrl,
			},
		}

		data = append(data, datum)
	}

	return dto.PostRepliesPaginationResponse{
		Data: dto.PostResponse{
			ID:       post.ID,
			Text:     post.Text,
			ParentID: post.ParentID,
			User: dto.UserResponse{
				ID:       post.UserID.String(),
				Name:     post.User.Name,
				Bio:      post.User.Bio,
				UserName: post.User.Username,
				ImageUrl: post.User.ImageUrl,
			},
			Replies: data,
		},
		PaginationResponse: dto.PaginationResponse{
			Page:    replies.Page,
			PerPage: replies.PerPage,
			MaxPage: replies.MaxPage,
			Count:   replies.Count,
		},
	}, nil
}

func (s *postService) DeletePostById(ctx context.Context, postId uint64) error {
	_, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.ErrGetPostById
	}

	if err := s.postRepo.DeletePostById(ctx, nil, postId); err != nil {
		return dto.ErrDeletePostById
	}

	return nil
}

func (s *postService) UpdatePostById(ctx context.Context, userId string, postId uint64, req dto.PostUpdateRequest) (dto.PostResponse, error) {
	post, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.PostResponse{}, dto.ErrGetPostById
	}

	if post.UserID.String() != userId {
		return dto.PostResponse{}, dto.ErrUnauthorized
	}

	post.Text = req.Text

	result, err := s.postRepo.UpdatePostById(ctx, nil, postId, post)
	if err != nil {
		return dto.PostResponse{}, dto.ErrUpdatePostById
	}

	return dto.PostResponse{
		ID:       result.ID,
		Text:     result.Text,
		ParentID: result.ParentID,
		User: dto.UserResponse{
			ID:       result.UserID.String(),
			Name:     result.User.Name,
			Bio:      result.User.Bio,
			UserName: result.User.Username,
			ImageUrl: result.User.ImageUrl,
		},
	}, nil
}

func (s *postService) GetAllPosts(ctx context.Context, req dto.PaginationRequest) (dto.PostPaginationResponse, error) {
	dataWithPaginate, err := s.postRepo.GetAllPostsWithPagination(ctx, nil, req)
	if err != nil {
		return dto.PostPaginationResponse{}, err
	}

	var data []dto.PostResponse
	for _, post := range dataWithPaginate.Posts {
		datum := dto.PostResponse{
			ID:       post.ID,
			Text:     post.Text,
			ParentID: post.ParentID,
			User: dto.UserResponse{
				ID:       post.UserID.String(),
				Name:     post.User.Name,
				Bio:      post.User.Bio,
				UserName: post.User.Username,
				ImageUrl: post.User.ImageUrl,
			},
		}

		data = append(data, datum)
	}

	return dto.PostPaginationResponse{
		Data: data,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
