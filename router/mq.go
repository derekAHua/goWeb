package router

import (
	"fmt"
	"github.com/derekAHua/goLib/rmq"
	"github.com/gin-gonic/gin"
)

const (
	tag = "YZH_tag"
)

// MQ mq消费者回调入口
func MQ(g *gin.Engine) {
	// rocketMQ 消费回调handler注册 , service 需要在 helpers/init.go 中注册（InitRocketMq中的 rmq.InitRmq）。
	// 一个应用尽可能用一个Topic，而消息子类型则可以用tags来标识。tags可以由应用自由设置，
	// 只有生产者在发送消息设置了tags，消费方在订阅消息时才可以利用tags通过broker做消息过滤
	rmq.Use(g, "demo", []string{
		tag, // tags
	}, func(ctx *gin.Context, msg rmq.Message) error {
		fmt.Println("msgContent:", string(msg.GetContent()))
		fmt.Println("msgOffsetId:", msg.GetID())
		return nil
	})
}
