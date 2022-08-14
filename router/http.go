package router

import (
	"github.com/derekAHua/goLib/middleware"
	"goWeb/controller"

	"github.com/gin-gonic/gin"
)

func Http(engine *gin.Engine) {
	router := engine.Group("/web")

	router.Use(middleware.Recover)

	testGroup := router.Group("/user")
	{
		testGroup.GET("/getUserInfo", controller.User.GetUserInfo)
	}

}
