package models

import "github.com/derekAHua/goLib/model"

// User 用户表
type User struct {
	model.Model
	UserId      uint64 `gorm:"column:user_id;unique;default:0;not null" json:"userId"` // 用户唯一Id
	RealName    string `gorm:"column:real_name;not null" json:"realName"`              // 真实姓名
	FirstLetter string `gorm:"column:first_letter;not null" json:"firstLetter"`        // 首字母，用于排序
	Phone       string `gorm:"column:phone;not null" json:"phone"`                     // 手机号
	Email       string `gorm:"column:email;not null" json:"email"`                     // email
	Tags        string `gorm:"column:tags;not null" json:"tags"`                       // 标签
	Status      int8   `gorm:"column:status;default:0;not null" json:"status"`         // 0:有效
}

func (*User) TableName() string {
	return "tblUser"
}
