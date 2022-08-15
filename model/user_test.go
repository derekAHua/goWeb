package model

import (
	"github.com/gin-gonic/gin"
	"goWeb/conf"
	"goWeb/test"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/8/15 22:17
// @Version 1.0

func TestNewUserWithTx(t *testing.T) {

	test.Init()
	//conf.Init()
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
