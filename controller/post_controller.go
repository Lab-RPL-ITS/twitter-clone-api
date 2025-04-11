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
	PostController interface {
		CreatePost(ctx *gin.Context)
		GetPostById(ctx *gin.Context)
		DeletePostById(ctx *gin.Context)
		UpdatePostById(ctx *gin.Context)
		GetAllPosts(ctx *gin.Context)
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
	postIdStr := ctx.Param("post_id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
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

	res := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_POST_BY_ID,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *postController) DeletePostById(ctx *gin.Context) {
	postIdStr := ctx.Param("post_id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.postService.DeletePostById(ctx.Request.Context(), postId)

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_POST, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *postController) UpdatePostById(ctx *gin.Context) {
	var post dto.PostUpdateRequest
	userId := ctx.GetString("user_id")

	if err := ctx.ShouldBind(&post); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	postIdStr := ctx.Param("post_id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.postService.UpdatePostById(ctx.Request.Context(), userId, postId, post)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_POST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_POST, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *postController) GetAllPosts(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	posts, err := c.postService.GetAllPosts(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_POSTS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_ALL_POSTS,
		Data:    posts.Data,
		Meta:    posts.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
