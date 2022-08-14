package test

import (
	"github.com/derekAHua/goLib/env"
	"github.com/derekAHua/goLib/utils"
	"goWeb/conf"
)

// @Author: Derek
// @Description: For test.
// @Date: 2022/5/11 20:18
// @Version 1.0

var initialized bool

// Init test env.
func Init() {
	if initialized {
		return
	}
	initialized = true

	dir := utils.GetSourcePath()
	env.SetAppName("testing")
	env.SetRootPath(dir + "/../.")
	conf.Init()

	conf.Config.Log.Stdout = true
}
