package gorm

import (
	"errors"
	"gorm.io/gorm"
	"Zchat/internal/dao"
	"Zchat/internal/dto/request"
	"Zchat/internal/model"
	"Zchat/pkg/zlog"
)
type groupInfoService struct {

}

var GroupInfoService = new(groupInfoService)

// SaveGroup 保存群聊，创建群聊也用这接口
func (g *groupInfoService) SaveGroup(groupReq request.GroupRequest) error {
	var group model.GroupInfo
	res := dao.GormDB.First(&group, "uuid = ?", groupReq.Uuid)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// 创建群聊

		} else {
			zlog.Error(res.Error.Error())
			return res.Error
		}
	}
	return nil
}

// GetAllMembers 获取所有成员信息
func (g *groupInfoService) GetAllMembers(groupId string) ([]model.UserInfo, error) {
	var group model.GroupInfo
	res := dao.GormDB.Preload("Members").First(&group, "uuid = ?", groupId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			zlog.Info("group not existed")
			return nil, nil
		} else {
			zlog.Error(res.Error.Error())
			return nil, res.Error
		}
	} else {
		return group.Members, nil
	}
}