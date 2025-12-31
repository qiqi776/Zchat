package model

import (
	"gorm.io/gorm"
	"time"
)

type GroupInfo struct {
	Id        int64          `gorm:"column:id;primaryKey;comment:自增id" json:"id"`
	Uuid      string         `gorm:"column:uuid;uniqueIndex;type:char(20);not null;comment:群组唯一id" json:"uuid"`
	Name      string         `gorm:"column:name;type:varchar(20);not null;comment:群名称" json:"name"`
	Notice    string         `gorm:"column:notice;type:varchar(500);comment:群公告" json:"notice"`
	Members   []string       `gorm:"column:members;many2many:user_groups;comment:群组成员" json:"members"`
	MemberCnt int            `gorm:"column:member_cnt;default:1;comment:群人数" json:"memberCnt"`
	OwnerId   string         `gorm:"column:owner_id;type:char(20);not null;comment:群主uuid" json:"ownerId"`
	AddMode   bool           `gorm:"column:add_mode;type:tinyint(1);default:false;comment:加群方式，0.直接，1.审核" json:"addMode"`
	Avatar    string         `gorm:"column:avatar;type:char(255);default:default_avatar.png;not null;comment:头像" json:"avatar"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"-"` // json:"-" 表示不返回给前端
}

func (GroupInfo) TableName() string {
	return "group_info"
}