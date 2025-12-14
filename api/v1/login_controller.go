package v1

import (
	"Zchat/internal/dto/request"
	"Zchat/internal/model"
	"Zchat/pkg/log"
	"net/http"
	"time"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
)

// 注册
func Register(c *gin.Context) {
	var registerReq request.RegisterRequest
	if err := c.BindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, u := range model.UserTable {
		if u.TelePhone == registerReq.Telephone {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user has exists in system"})
			return
		}
	}
	// 100000000000 - 99999999999 11位随机数
	tmpInt1 := rand.Intn(10000000000)
	tmpInt2 := (rand.Intn(9) + 1) * 10000000000
	tmpInt := tmpInt1 + tmpInt2
	newUserInfo := model.UserInfo{
		Id:        "U" + strconv.Itoa(tmpInt),
		TelePhone: registerReq.Telephone,
		Password:  registerReq.Password,
		NickName:  registerReq.NickName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	model.UserTable = append(model.UserTable, newUserInfo)
	// TODO 数据库插入

	log.GetZapLogger().Info(*util.ToJSONString(newUserInfo))
	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

// 登录
func Login(c *gin.Context) {
	var loginReq request.Loginrequest
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, u := range model.UserTable {
		if u.TelePhone == loginReq.Telephone && u.Password == loginReq.Password {
			c.JSON(http.StatusOK, gin.H{"message": "login success"})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "telephone or password incorrect"})
}