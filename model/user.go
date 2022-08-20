package model

import (
	"github.com/derekAHua/goLib/errors"
	"github.com/derekAHua/goLib/model"
	"github.com/derekAHua/goLib/zlog"
	"github.com/gin-gonic/gin"
	"goWeb/conf"
	"goWeb/models"
	"gorm.io/gorm"
)

var (
	errGetUserInfo = errors.Err{ErrNo: 100, ErrMsg: "获取用户信息失败"}
)

type (
	User interface {
		GetUSerById(userId uint64) (defaultUser models.User, err error)
		UpdateStatus(userId uint64) (err error)
	}

	user struct {
		ctx *gin.Context
		model.BaseModel
	}
)

func (u *user) UpdateStatus(userId uint64) (err error) {
	m := map[string]interface{}{"status": 1}
	_, err = u.Update(m, model.Where("user_id = ?", userId))
	return
}

func (u *user) GetUSerById(userId uint64) (defaultUser models.User, err error) {
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
