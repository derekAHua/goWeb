package model

import (
	"github.com/gin-gonic/gin"
	"goWeb/conf"
	"goWeb/models"
	"goWeb/test"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/8/15 22:17
// @Version 1.0

func TestSaveOrCreate(t *testing.T) {
	test.Init()
	ctx := &gin.Context{}

	err := NewUser(ctx).SaveOrCreate(models.User{
		UserId:   1,
		RealName: "aaa",
		Phone:    "232",
		Email:    "a.com",
	})

	if err != nil {
		t.Log(err)
	}
}

func TestNewUserWithTx(t *testing.T) {

	test.Init()
	ctx := &gin.Context{}

	var err error
	tx := conf.MysqlClientTest.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = NewUserWithTx(ctx, tx).UpdateStatus(1)
	if err != nil {
		t.Error(err)
		return
	}

	err = NewUserWithTx(ctx, tx).UpdateStatus(2)
	//err = errors.New("err")
	if err != nil {
		t.Error(err)
		return
	}
}
