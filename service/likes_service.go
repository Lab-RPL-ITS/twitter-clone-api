package service

import (
	"context"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
)

type (
	LikesService interface {
		LikePostById(ctx context.Context, postId uint64, userId string) error
		UnLikePostById(ctx context.Context, postId uint64, userId string) error
	}

	likesService struct {
		likesRepo  repository.LikesRepository
		postRepo   repository.PostRepository
		jwtService JWTService
	}
)

func NewLikesService(likesRepo repository.LikesRepository, postRepo repository.PostRepository, jwtService JWTService) LikesService {
	return &likesService{
		likesRepo:  likesRepo,
		postRepo:   postRepo,
		jwtService: jwtService,
	}
}

func (s *likesService) LikePostById(ctx context.Context, postId uint64, userId string) error {
	_, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.ErrGetPostById
	}

	err = s.likesRepo.LikePostById(ctx, nil, postId, userId)
	if err != nil {
		return dto.ErrLikePostById
	}

	err = s.postRepo.UpdateLikesCount(ctx, nil, postId, 1)
	if err != nil {
		return dto.ErrLikePostById
	}

	return nil
}

func (s *likesService) UnLikePostById(ctx context.Context, postId uint64, userId string) error {
	err := s.likesRepo.CheckLikedPost(ctx, nil, postId, userId)
	if err != nil {
		return dto.ErrCheckLikedPost
	}

	err = s.likesRepo.ErrUnlikePostById(ctx, nil, postId, userId)
	if err != nil {
		return dto.ErrUnlikePostById
	}

	err = s.postRepo.UpdateLikesCount(ctx, nil, postId, -1)
	if err != nil {
		return dto.ErrUnlikePostById
	}

	return nil
}
