package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	Id            int64          `gorm:"column:id;primaryKey;comment:自增id" json:"id"`
	Uuid          string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);comment:用户唯一id" json:"uuid"`
	NickName      string         `gorm:"column:nickname;type:varchar(20);not null;comment:昵称" json:"nickname"`
	TelePhone     string         `gorm:"column:telephone;uniqueIndex;not null;type:char(11);comment:电话" json:"telephone"`
	Email         string         `gorm:"column:email;type:char(30);comment:邮箱" json:"email"`
	Avatar        string         `gorm:"column:avatar;type:char(255);default:default_avatar.png;not null;comment:头像" json:"avatar"`
	Gender        bool           `gorm:"column:gender;type:tinyint(1);comment:性别" json:"gender"`
	Signature     string         `gorm:"column:signature;type:varchar(50);comment:个性签名" json:"signature"`
	Password      string         `gorm:"column:password;type:char(18);not null;comment:密码" json:"-"` // 安全起见，密码永远不要传给前端
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"createdAt"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"-"` // 不返回
	LastOnlineAt  time.Time      `gorm:"column:last_online_at;type:datetime;comment:上次登录时间" json:"lastOnlineAt"`
	LastOfflineAt time.Time      `gorm:"column:last_offline_at;type:datetime;comment:最近离线时间" json:"lastOfflineAt"`
	IsAdmin       bool           `gorm:"column:is_admin;type:tinyint(1);not null;comment:是否是管理员，0.不是，1.是" json:"isAdmin"`
}

func (UserInfo) TableName() string {
	return "user_info"
}