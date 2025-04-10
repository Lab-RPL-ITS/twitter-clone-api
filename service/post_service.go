package service

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
)

type (
	PostService interface {
		CreatePost(ctx context.Context, req dto.PostCreateRequest) (dto.PostResponse, error)
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

func (s *postService) CreatePost(ctx context.Context, req dto.PostCreateRequest) (dto.PostResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, req.UserID)

	if err != nil {
		return dto.PostResponse{}, dto.ErrGetUserById
	}

	if req.ParentID != nil {
		_, err := s.postRepo.GetPostById(ctx, nil, *req.ParentID)
		if err != nil {
			return dto.PostResponse{}, dto.ErrGetPostById
		}
	}

	if err != nil {
		return dto.PostResponse{}, dto.ErrParseParentID
	}

	post := entity.Post{
		Text:     req.Text,
		UserID:   user.ID,
		ParentID: req.ParentID,
	}

	result, err := s.postRepo.CreatePost(ctx, nil, post)
	if err != nil {
		return dto.PostResponse{}, dto.ErrCreatePost
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
