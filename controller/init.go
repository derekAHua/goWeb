package controller

// @Author: Derek
// @Description: Controllers
// @Date: 2022/8/14 22:24
// @Version 1.0

var (
	User user
)

type (
	user struct{}
)

func init() {
	User = user{}
}
