package dto

import (
	"errors"

	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_GET_POST_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_POST             = "failed create post"
	MESSAGE_FAILED_GET_POST_ID             = "failed get post id"
	MESSAGE_FAILED_UPDATE_POST             = "failed update post"

	// Succcess
	MESSAGE_SUCCESS_CREATE_POST    = "success create post"
	MESSAGE_SUCCESS_GET_POST_BY_ID = "success get post by id"
	MESSAGE_SUCCESS_DELETE_POST    = "success delete post"
	MESSAGE_SUCCESS_UPDATE_POST    = "success update post"
)

var (
	ErrCreatePost     = errors.New("failed to create post")
	ErrGetPostById    = errors.New("post not found")
	ErrParseParentID  = errors.New("failed to parse parent id")
	ErrDeletePostById = errors.New("failed to delete post")
	ErrUpdatePostById = errors.New("failed to update post")
)

type (
	PostCreateRequest struct {
		Text     string     `json:"text" form:"text" binding:"required"`
		ParentID *uuid.UUID `json:"parent_id," form:"parent_id"`
	}

	PostResponse struct {
		ID       string       `json:"id"`
		Text     string       `json:"text"`
		ParentID *uuid.UUID   `json:"parent_id"`
		User     UserResponse `json:"user"`
	}

	PostUpdateRequest struct {
		Text string `json:"text" form:"text" binding:"required"`
	}
)
