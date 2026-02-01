package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Zchat/internal/dto/request"
	"Zchat/internal/service/gorm"
)

// CreateGroup 创建群聊
func CreateGroup(c *gin.Context) {
	var createGroupReq request.CreateGroupRequest
	if err := c.BindJSON(&createGroupReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	if err := gorm.GroupInfoService.CreateGroup(createGroupReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "create group success",
	})
}

// LoadMyGroup 获取我创建的群聊
func LoadMyGroup(c *gin.Context) {
	var loadGroupReq request.LoadMyGroupRequest
	if err := c.BindJSON(&loadGroupReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	groupList, err := gorm.GroupInfoService.LoadMyGroup(loadGroupReq.OwnerId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "load my group success",
		"data":    groupList,
	})
}