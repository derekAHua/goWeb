package model

import (
	"github.com/derekAHua/goLib/errors"
	"github.com/derekAHua/goLib/model"
	"github.com/derekAHua/goLib/zlog"
	"github.com/gin-gonic/gin"
	"goWeb/conf"
	"gorm.io/gorm"
)

var (
	errGetUserInfo = errors.Err{ErrNo: 100, ErrMsg: "获取用户信息失败"}
)

type (
	user interface {
		GetUSerById(userId uint64) (defaultUser TblUser, err error)
	}

	defaultUser struct {
		ctx *gin.Context
		model.BaseModel
	}
)

func (u *defaultUser) GetUSerById(userId uint64) (defaultUser TblUser, err error) {
	err = u.GetOne(&defaultUser, model.Where("user_id = ?", userId), model.UnDelete())
	if err != nil {
		zlog.ErrorF(u.ctx, errGetUserInfo.Error()+",userId=[%d],err=[%s]", userId, err.Error())
		err = errGetUserInfo
	}

	return
}

func NewUser(ctx *gin.Context) user {
	return &defaultUser{ctx, model.NewBaseModel(conf.MysqlClientTest.WithContext(ctx).Table("tblUser"))}
}

func NewUserWithTx(ctx *gin.Context, tx *gorm.DB) user {
	return &defaultUser{ctx, model.NewBaseModel(tx.Table("tblUser"))}
}
