package gorm

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"Zchat/internal/dao"
	"Zchat/internal/dto/respond"
	"Zchat/internal/model"
	"Zchat/pkg/zlog"
)

type userContactService struct {

}

var UserContactService = new(userContactService)

func (u *userContactService) GetUserList(ownerId string) (string, []respond.MyUserListRespond, error) {
	var userList []model.UserContact
	if res := dao.GormDB.Order("created_at DESC").Where("user_id = ?", ownerId).Find(&userList); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			message := "contact not found"
			zlog.Info(message)
			return "", nil, nil
		}
		zlog.Error(res.Error.Error())
		return "", nil, res.Error
	}
	
	log.Println(userList)
	var userListRsp []respond.MyUserListRespond
	for _, contact := range userList {
		contactId := contact.ContactId
		var user model.UserInfo
		if res := dao.GormDB.First(&user, "uuid = ?", contactId); res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				zlog.Error("contact not found in UserInfo table")
				return "", nil, res.Error
			} else {
				zlog.Error(res.Error.Error())
				return "", nil, res.Error
			}
		}
		userListRsp = append(userListRsp, respond.MyUserListRespond{
			UserId:   user.Uuid,
			UserName: user.Nickname,
			Avatar:   user.Avatar,
		})
	}
	return "", userListRsp, nil
}