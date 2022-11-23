package model

import (
	"github.com/renatozhang/goorm/id_gen"
	"gorm.io/gorm"
)

const (
	UserSexMan   = 1
	UserSexWomen = 2
)

type User struct {
	UserId   int64  `json:"user_id" gorm:"user_id"`
	NickName string `json:"nickname" gorm:"column:nickname"`
	Username string `json:"user" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
	Sex      int    `json:"sex" gorm:"sex"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := id_gen.GetID()
	if err != nil {
		return
	}
	u.UserId = int64(uid)
	return
}
