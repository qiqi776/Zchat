package model

import "time"

type UserContactApply struct {
	Id          int64     `gorm:"column:id;primaryKey;comment:申请id" json:"id"`
	UserId      string    `gorm:"column:user_id;index;type:char(20);not null;comment:申请人id" json:"userId"`
	ContactId   string    `gorm:"column:apply_contact_id;index;type:char(20);not null;comment:被申请人id" json:"contactId"`
	ContactType bool      `gorm:"column:apply_contact_type;type:tinyint(1);not null;comment:被申请人类型" json:"contactType"`
	Status      bool      `gorm:"column:status;not null;type:tinyint(1);comment:申请状态，0.申请中，1.申请通过" json:"status"`
	Message     string    `gorm:"column:message;type:varchar(100);comment:申请信息" json:"message"`
	LastApplyAt time.Time `gorm:"column:last_apply_time;type:datetime;not null;comment:最后申请时间" json:"lastApplyAt"`
}

func (UserContactApply) TableName() string {
	return "user_contact_apply"
}