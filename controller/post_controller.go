package controller

import (
	"net/http"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/Lab-RPL-ITS/twitter-clone-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	PostController interface {
		Create(ctx *gin.Context)
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

func (c *postController) Create(ctx *gin.Context) {
	var post dto.PostCreateRequest
	if err := ctx.ShouldBind(&post); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.postService.CreatePost(ctx.Request.Context(), post)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_POST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_POST, result)
	ctx.JSON(http.StatusOK, res)
}
