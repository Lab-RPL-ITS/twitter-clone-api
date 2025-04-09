package dto

import (
	"errors"
	"mime/multipart"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER      = "failed create user"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_GET_USER           = "failed get user"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success create user"
	MESSAGE_SUCCESS_GET_USER      = "success get user"
	MESSAGE_SUCCESS_LOGIN         = "success login"
)

var (
	ErrCreateUser            = errors.New("failed to create user")
	ErrGetUserById           = errors.New("failed to get user by id")
	ErrUsernameAlreadyExists = errors.New("username already exist")
	ErrUsernameNotFound      = errors.New("username not found")
	ErrPasswordNotMatch      = errors.New("password not match")
)

type (
	UserCreateRequest struct {
		Name     string                `json:"name" form:"name" binding:"required"`
		UserName string                `json:"username" form:"username" binding:"required"`
		Bio      string                `json:"bio" form:"bio" binding:"required"`
		Image    *multipart.FileHeader `json:"image" form:"image"`
		Password string                `json:"password" form:"password" binding:"required"`
	}

	UserResponse struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		UserName string  `json:"username"`
		Bio      string  `json:"bio"`
		ImageUrl *string `json:"image_url"`
	}

	UserLoginRequest struct {
		UserName string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}
)
