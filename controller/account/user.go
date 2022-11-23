package account

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/goorm/dal/db"
	"github.com/renatozhang/goorm/model"
	"github.com/renatozhang/gostudy/mercury/util"
)

func RegisterHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = db.Register(&user)
	if err == db.ErrUserExists {
		util.ResponseError(ctx, util.ErrCodeUserExist)
		return
	}
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(ctx, nil)
}

func LoginHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = db.Login(&user)
	if err == db.ErrUserNotExists {
		util.ResponseError(ctx, util.ErrCodeUserNotExists)
		return
	}
	if err == db.ErrUserPasswordWrong {
		util.ResponseError(ctx, util.ErrCodeUserPasswordWrong)
		return
	}
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(ctx, nil)
}
