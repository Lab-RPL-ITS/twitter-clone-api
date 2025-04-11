package controller

import (
	"net/http"
	"strconv"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/Lab-RPL-ITS/twitter-clone-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	LikesController interface {
		LikePostById(ctx *gin.Context)
		UnlikePostById(ctx *gin.Context)
	}

	likesController struct {
		likesService service.LikesService
	}
)

func NewLikesController(likesService service.LikesService) LikesController {
	return &likesController{
		likesService: likesService,
	}
}

func (c *likesController) LikePostById(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	userId := ctx.GetString("user_id")

	postIdUint, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.likesService.LikePostById(ctx, postIdUint, userId)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIKE_POST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LIKE_POST, nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *likesController) UnlikePostById(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	userId := ctx.GetString("user_id")

	postIdUint, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.likesService.UnLikePostById(ctx, postIdUint, userId)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UNLIKE_POST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UNLIKE_POST, nil)
	ctx.JSON(http.StatusOK, response)
}
