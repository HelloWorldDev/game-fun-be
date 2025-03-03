package api

import (
	"my-token-ai-be/internal/request"
	"my-token-ai-be/internal/response"
	"my-token-ai-be/internal/service"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
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
// @Router /users/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(http.StatusBadRequest, "Invalid request parameters", err))
		return
	}
	res := u.userService.Login(req)
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
// @Router /users/my_info [get]
func (u *UserHandler) MyInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Err(http.StatusUnauthorized, "Address not found in context", nil))
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, response.Err(http.StatusUnauthorized, "Invalid address type in context", nil))
		return
	}
	res := u.userService.MyInfo(userIDStr)
	c.JSON(res.Code, res)
}
