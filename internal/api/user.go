package api

import (
	"my-token-ai-be/internal/request"
	"my-token-ai-be/internal/response"
	"my-token-ai-be/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserServiceImpl
}

func NewUserHandler(userService *service.UserServiceImpl) *UserHandler {
	return &UserHandler{userService: userService}
}

// Login 用户登录
// @Summary 用户钱包登录
// @Description 通过钱包地址和签名进行登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param login body request.LoginRequest true "登录请求参数"
// @Success 200 {object} response.LoginResponse "登录成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users/{chain_type}/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(http.StatusBadRequest, "Invalid request parameters", err))
		return
	}
	chainType, errResp := ParseChainTypeWithResponse(c)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}
	res := u.userService.Login(req, chainType)
	c.JSON(res.Code, res)
}

// MyInfo 获取用户信息
// @Summary 获取当前用户信息
// @Description 根据 JWT Token 获取当前用户的详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=response.MyInfoResponse} "成功返回用户信息"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users/{chain_type}/my_info [get]
func (u *UserHandler) MyInfo(c *gin.Context) {
	userID, errResp := GetUserIDFromContext(c)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}
	chainType, errResp := ParseChainTypeWithResponse(c)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}
	res := u.userService.MyInfo(userID, chainType)
	c.JSON(res.Code, res)
}

// InviteCode 获取用户邀请码信息
// @Summary 获取用户邀请码信息
// @Description 根据用户 ID 获取用户的邀请码和邀请数量
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=response.InviteCodeResponse} "成功返回用户邀请码信息"
// @Failure 401 {object} response.Response "未授权"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users/{chain_type}/invite/code [get]
func (u *UserHandler) InviteCode(c *gin.Context) {
	userID, errResp := GetUserIDFromContext(c)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}
	chainType, errResp := ParseChainTypeWithResponse(c)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}
	res := u.userService.GetCode(userID, chainType)
	c.JSON(res.Code, res)
}
