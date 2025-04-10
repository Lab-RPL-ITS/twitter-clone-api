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
		GetPostById(ctx context.Context, postId uuid.UUID) (dto.PostResponse, error)
		DeletePostById(ctx context.Context, postId uuid.UUID) error
		UpdatePostById(ctx context.Context, userId string, postId uuid.UUID, req dto.PostUpdateRequest) (dto.PostResponse, error)
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
		ID:       result.ID.String(),
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

func (s *postService) GetPostById(ctx context.Context, postId uuid.UUID) (dto.PostResponse, error) {
	post, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.PostResponse{}, dto.ErrGetPostById
	}

	return dto.PostResponse{
		ID:       post.ID.String(),
		Text:     post.Text,
		ParentID: post.ParentID,
		User: dto.UserResponse{
			ID:       post.UserID.String(),
			Name:     post.User.Name,
			Bio:      post.User.Bio,
			UserName: post.User.Username,
			ImageUrl: post.User.ImageUrl,
		},
	}, nil
}

func (s *postService) DeletePostById(ctx context.Context, postId uuid.UUID) error {
	_, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.ErrGetPostById
	}

	if err := s.postRepo.DeletePostById(ctx, nil, postId); err != nil {
		return dto.ErrDeletePostById
	}

	return nil
}

func (s *postService) UpdatePostById(ctx context.Context, userId string, postId uuid.UUID, req dto.PostUpdateRequest) (dto.PostResponse, error) {
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
		ID:       result.ID.String(),
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
