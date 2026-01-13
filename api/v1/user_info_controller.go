package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Zchat/internal/dto/request"
	"Zchat/internal/service/gorm"
)

// 注册
func Register(c *gin.Context) {
	var registerReq request.RegisterRequest
	if err := c.BindJSON(&registerReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	message, userInfoStr, err := gorm.UserInfoService.Register(c, registerReq)
	if message != "" && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": message,
		})
		return
	} else if message == "" && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	} else if message == "" && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "register success",
			"data":    userInfoStr,
		})
		return
	}
}

// 登录
func Login(c *gin.Context) {
	var loginReq request.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	message, userInfoStr, err := gorm.UserInfoService.Login(c, loginReq)
	if message != "" && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": message,
		})
		return
	} else if message == "" && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	} else if message == "" && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "login success",
			"data":    userInfoStr,
		})
		return
	}
}
