package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_LIKE_POST   = "failed like post"
	MESSAGE_FAILED_UNLIKE_POST = "failed unlike post"

	// Succcess
	MESSAGE_SUCCESS_LIKE_POST   = "success like post"
	MESSAGE_SUCCESS_UNLIKE_POST = "success unlike post"
)

var (
	ErrLikePostById   = errors.New("failed to like post")
	ErrCheckLikedPost = errors.New("failed to check liked post")
	ErrUnlikePostById = errors.New("failed to unlike post")
)

type (
	LikesRequest struct {
		PostID uint64 `json:"post_id" form:"post_id" binding:"required"`
	}
)
