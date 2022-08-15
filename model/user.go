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
	User interface {
		GetUSerById(userId uint64) (defaultUser TblUser, err error)
	}

	user struct {
		ctx *gin.Context
		model.BaseModel
	}
)

func (u *user) GetUSerById(userId uint64) (defaultUser TblUser, err error) {
	err = u.GetOne(&defaultUser, model.Where("user_id = ?", userId), model.UnDelete())
	if err != nil {
		zlog.ErrorF(u.ctx, errGetUserInfo.Error()+",userId=[%d],err=[%s]", userId, err.Error())
		err = errGetUserInfo
	}

	return
}

func NewUser(ctx *gin.Context) User {
	return &user{ctx, model.NewBaseModel(conf.MysqlClientTest.WithContext(ctx).Table("tblUser"))}
}

func NewUserWithTx(ctx *gin.Context, tx *gorm.DB) User {
	return &user{ctx, model.NewBaseModel(tx.Table("tblUser"))}
}
