package routes

import (
	"network_framework/core/network"
	"network_framework/app/controllers/web"
)

//请参考gin 用法
func RegistersWeb() {
	//HelloWorld
	network.Wroutes.GET("/",web.HelloWorld)
	//open websocket
	network.Wroutes.GET("/socket.io", network.WebsocketIo)
}