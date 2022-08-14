package conf

import (
	"github.com/derekAHua/goLib/conf"
	"github.com/derekAHua/goLib/env"
	"github.com/derekAHua/goLib/redis"
	"github.com/derekAHua/goLib/zlog"
	"gorm.io/gorm"
)

var (
	Config   conf.ConfigConf
	Api      ApiConf
	Resource conf.ResourceConf
	App      AppConf
)

var (
	MysqlClientTest *gorm.DB
	RedisClient     *redis.Redis
)

func Init() {
	env.SetAppName("goWeb")
	conf.InitConf(&Config, &Api, &Resource, &App)
	zlog.Init(Config.Log)

	mMysql, mRedis := conf.InitResource(&Resource)

	for name := range Resource.Mysql {
		switch name {
		case "test":
			MysqlClientTest = mMysql[name]
		default:
		}
	}

	for name := range Resource.Redis {
		switch name {
		case "demo":
			RedisClient = mRedis[name]
		default:
		}
	}
}

func Clear() {
	zlog.CloseLogger()
	MysqlClientTest = nil
	_ = RedisClient.Close()
}
