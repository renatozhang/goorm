package db

import (
	"github.com/renatozhang/goorm/app"
	"github.com/renatozhang/goorm/model"
	"github.com/renatozhang/goorm/util"
)

const (
	PasswordSalt = "tLcZlok88secbtcU5tJoJ&KQdUu9$&vR"
)

/*
	UserId   int64  `json:"user_id" gorm:"user_id"`
	NickName string `json:"nickname" gorm:"nickname"`
	Username string `json:"user" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
	Sex      int    `json:"sex" gorm:"sex"`
*/

func Register(user *model.User) (err error) {
	var count int64
	err = app.MySQL.DB().Debug().Select("UserId").First(&user).Where("username = ?", &user.Username).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		err = ErrUserExists
		return
	}
	pwd := user.Password + PasswordSalt
	user.Password = util.Md5([]byte(pwd))
	err = app.MySQL.DB().Debug().Select("UserId", "NickName", "Username", "Password", "Email", "Sex").Create(&user).Error
	return
}
func Login(user *model.User) (err error) {
	originPasswd := user.Password
	find := app.MySQL.DB().Debug().Select("UserId", "Username", "Password").Where("username = ?", user.Username).First(user)
	if find.RowsAffected < 1 {
		err = ErrUserNotExists
		return
	}
	passwd := originPasswd + PasswordSalt
	originPasswordSalt := util.Md5([]byte(passwd))
	if originPasswordSalt != user.Password {
		err = ErrUserPasswordWrong
		return
	}
	return
}
