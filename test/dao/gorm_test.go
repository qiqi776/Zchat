package dao

import (
	"strconv"
	"testing"
	"time"
	"Zchat/internal/dao"
	"Zchat/internal/model"
	"Zchat/pkg/util/random"
)

func TestCreate(t *testing.T) {
	userInfo := &model.UserInfo{
		Uuid:      "U" + strconv.Itoa(random.GetRandomInt(11)),
		Nickname:  "apylee",
		Telephone: "180323532112",
		Email:     "1212312312@qq.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		IsAdmin:   true,
	}
	_ = dao.GormDB.Create(userInfo)
}