package service

import (
	"context"
	"fmt"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
	"github.com/Lab-RPL-ITS/twitter-clone-api/helpers"
	"github.com/Lab-RPL-ITS/twitter-clone-api/repository"
	"github.com/Lab-RPL-ITS/twitter-clone-api/utils"
	"github.com/google/uuid"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		GetUserByUsername(ctx context.Context, username string) (dto.UserResponse, error)
		UpdateUser(ctx context.Context, userId string, req dto.UserProfileUpdateRequest) (dto.UserResponse, error)
		GetUserPosts(ctx context.Context, username string, req dto.UserPostsPaginationRequest) (dto.PostPaginationResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		postRepo   repository.PostRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, postRepo repository.PostRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		postRepo:   postRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	_, flag, _ := s.userRepo.CheckUsername(ctx, nil, req.UserName)
	if flag {
		return dto.UserResponse{}, dto.ErrUsernameAlreadyExists
	}

	user := entity.User{
		Name:     req.Name,
		Username: req.UserName,
		ImageUrl: nil,
		Bio:      nil,
		Password: req.Password,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:       userReg.ID.String(),
		Name:     userReg.Name,
		UserName: userReg.Username,
		Bio:      userReg.Bio,
		ImageUrl: userReg.ImageUrl,
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:       user.ID.String(),
		Name:     user.Name,
		UserName: user.Username,
		Bio:      user.Bio,
		ImageUrl: user.ImageUrl,
	}, nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckUsername(ctx, nil, req.UserName)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrUsernameNotFound
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	token := s.jwtService.GenerateToken(check.ID.String())

	return dto.UserLoginResponse{
		Token: token,
	}, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (dto.UserResponse, error) {
	user, _, err := s.userRepo.CheckUsername(ctx, nil, username)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUsernameNotFound
	}

	return dto.UserResponse{
		ID:       user.ID.String(),
		Name:     user.Name,
		UserName: user.Username,
		Bio:      user.Bio,
		ImageUrl: user.ImageUrl,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, userId string, req dto.UserProfileUpdateRequest) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	previousImage := user.ImageUrl

	if req.Image != nil {
		var filenamePtr *string
		imageId := uuid.New()
		ext := utils.GetExtensions(req.Image.Filename)

		filename := fmt.Sprintf("profile/%s.%s", imageId, ext)
		if err := utils.UploadFile(req.Image, filename); err != nil {
			return dto.UserResponse{}, err
		}
		filenamePtr = &filename
		user.ImageUrl = filenamePtr
	}

	if req.Bio != "" {
		var bioPtr *string
		bio := req.Bio
		bioPtr = &bio
		user.Bio = bioPtr
	}

	userUpdate, err := s.userRepo.UpdateUser(ctx, nil, userId, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	if previousImage != nil {
		if err := utils.DeleteFile(*previousImage); err != nil {
			return dto.UserResponse{}, err
		}
	}

	return dto.UserResponse{
		ID:       userUpdate.ID.String(),
		Name:     userUpdate.Name,
		UserName: userUpdate.Username,
		Bio:      userUpdate.Bio,
		ImageUrl: userUpdate.ImageUrl,
	}, nil
}

func (s *userService) GetUserPosts(ctx context.Context, username string, req dto.UserPostsPaginationRequest) (dto.PostPaginationResponse, error) {
	dataWithPaginate, err := s.postRepo.GetAllPostsWithPaginationByUsername(ctx, nil, username, req)
	if err != nil {
		return dto.PostPaginationResponse{}, err
	}

	var data []dto.PostResponse
	for _, post := range dataWithPaginate.Posts {
		datum := dto.PostResponse{
			ID:         post.ID,
			Text:       post.Text,
			TotalLikes: post.TotalLikes,
			IsDeleted:  post.DeletedAt.Valid,
			ParentID:   post.ParentID,
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
