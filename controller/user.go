package controller

import (
	"github.com/derekAHua/goLib/base"
	"github.com/derekAHua/goLib/utils"
	"github.com/gin-gonic/gin"
	"goWeb/service"
)

// @Author: Derek
// @Description: User Controller.
// @Date: 2022/8/14 13:31
// @Version 1.0

type user struct{}

var User = user{}

func (user) GetUserInfo(ctx *gin.Context) {
	param := struct {
		UserId uint64 `form:"userId" binding:"required" label:"用户Id"`
	}{}

	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		base.RenderJsonFail(ctx, utils.Translate(err))
		return
	}

	user, err := service.User.GetUserInfo(ctx, param.UserId)
	if err != nil {
		base.RenderJsonFail(ctx, err)
		return
	}
	base.RenderJsonSuc(ctx, user)
	return
}
