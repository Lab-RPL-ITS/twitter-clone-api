package dto

import (
	"errors"

	"github.com/Lab-RPL-ITS/twitter-clone-api/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_POST_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_POST             = "failed create post"
	MESSAGE_FAILED_GET_POST_ID             = "failed get post id"
	MESSAGE_FAILED_UPDATE_POST             = "failed update post"
	MESSAGE_FAILED_GET_ALL_POSTS           = "failed get all posts"

	// Succcess
	MESSAGE_SUCCESS_CREATE_POST    = "success create post"
	MESSAGE_SUCCESS_GET_POST_BY_ID = "success get post by id"
	MESSAGE_SUCCESS_DELETE_POST    = "success delete post"
	MESSAGE_SUCCESS_UPDATE_POST    = "success update post"
	MESSAGE_SUCCESS_GET_ALL_POSTS  = "success get all posts"
)

var (
	ErrCreatePost     = errors.New("failed to create post")
	ErrGetPostById    = errors.New("post not found")
	ErrGetPostReplies = errors.New("failed to get post replies")
	ErrParseParentID  = errors.New("failed to parse parent id")
	ErrDeletePostById = errors.New("failed to delete post")
	ErrUpdatePostById = errors.New("failed to update post")
)

type (
	PostCreateRequest struct {
		Text     string  `json:"text" form:"text" binding:"required"`
		ParentID *uint64 `json:"parent_id," form:"parent_id"`
	}

	PostResponse struct {
		ID         uint64       `json:"id"`
		Text       string       `json:"text"`
		TotalLikes uint64       `json:"total_likes"`
		ParentID   *uint64      `json:"parent_id"`
		IsDeleted  bool         `json:"is_deleted"`
		User       UserResponse `json:"user"`
	}

	PostWithRepliesResponse struct {
		PostResponse
		Replies []PostResponse `json:"replies"`
	}

	PostUpdateRequest struct {
		Text string `json:"text" form:"text" binding:"required"`
	}

	PostPaginationResponse struct {
		Data []PostResponse `json:"data"`
		PaginationResponse
	}

	PostRepliesPaginationResponse struct {
		Data PostWithRepliesResponse `json:"data"`
		PaginationResponse
	}

	GetAllPostsRepositoryResponse struct {
		Posts []entity.Post `json:"posts"`
		PaginationResponse
	}

	GetAllRepliesRepositoryResponse struct {
		Replies []entity.Post `json:"replies"`
		PaginationResponse
	}
)
