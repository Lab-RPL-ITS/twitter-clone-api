package controller

import (
	"net/http"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/Lab-RPL-ITS/twitter-clone-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	PostController interface {
		CreatePost(ctx *gin.Context)
		GetPostById(ctx *gin.Context)
		DeletePostById(ctx *gin.Context)
	}

	postController struct {
		postService service.PostService
	}
)

func NewPostController(ps service.PostService) PostController {
	return &postController{
		postService: ps,
	}
}

func (c *postController) CreatePost(ctx *gin.Context) {
	var post dto.PostCreateRequest
	userId := ctx.GetString("user_id")

	if err := ctx.ShouldBind(&post); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.postService.CreatePost(ctx.Request.Context(), userId, post)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_POST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_POST, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *postController) GetPostById(ctx *gin.Context) {
	postId, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.postService.GetPostById(ctx.Request.Context(), postId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_POST_BY_ID, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *postController) DeletePostById(ctx *gin.Context) {
	postId, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.postService.DeletePostById(ctx.Request.Context(), postId)

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_POST, nil)
	ctx.JSON(http.StatusOK, res)
}
