package main

import (
	"github.com/derekAHua/goLib/generator"
	"github.com/derekAHua/irpc/server"
	"goWeb/conf"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/5/12 23:02
// @Version 1.0

func TestServer(t *testing.T) {
	server.New()
}

func TestGenerateModel(t *testing.T) {
	generator.GenerateModel("./model", "user")
}

func TestGenerateTable(t *testing.T) {
	conf.Init()
	generator.GenerateTable("./model", conf.MysqlClientTest, "test", "tblUser")
}
