package controller

import (
	"net/http"

	"github.com/Lab-RPL-ITS/twitter-clone-api/dto"
	"github.com/Lab-RPL-ITS/twitter-clone-api/service"
	"github.com/Lab-RPL-ITS/twitter-clone-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
		GetUserByUsername(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)
		CheckUsername(ctx *gin.Context)
		GetUserPosts(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Register(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Me(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	result, err := c.userService.GetUserById(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.userService.Verify(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, dto.ErrUsernameNotFound.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetUserByUsername(ctx.Request.Context(), username)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)
	var user dto.UserProfileUpdateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.UpdateUser(ctx.Request.Context(), userId, user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) CheckUsername(ctx *gin.Context) {
	var username dto.CheckUsernameRequest
	if err := ctx.ShouldBind(&username); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	_, err := c.userService.GetUserByUsername(ctx.Request.Context(), username.UserName)
	if err != nil && err.Error() == dto.ErrUsernameNotFound.Error() {
		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_USERNAME_AVAILABLE, nil)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_USERNAME_EXISTS, dto.ErrUsernameAlreadyExists.Error(), nil)
	ctx.JSON(http.StatusBadRequest, res)
}

func (c *userController) GetUserPosts(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, dto.ErrUsernameNotFound.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.UserPostsPaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_POST_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetUserPosts(ctx.Request.Context(), username, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_POSTS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_ALL_POSTS,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
