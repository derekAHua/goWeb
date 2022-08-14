package service

import (
	"github.com/gin-gonic/gin"
	"goWeb/model"
)

// @Author: Derek
// @Description: user Service.
// @Date: 2022/8/14 13:48
// @Version 1.0

type (
	user interface {
		// GetUserInfo return model.TblUser by userId.
		GetUserInfo(ctx *gin.Context, userId uint64) (user model.TblUser, err error)
	}

	defaultUser struct{}
)

func (u *defaultUser) GetUserInfo(ctx *gin.Context, userId uint64) (user model.TblUser, err error) {
	return model.NewUser(ctx).GetUSerById(userId)
}

func newUser() user {
	return &defaultUser{}
}