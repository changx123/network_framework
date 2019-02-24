package routes

import "blog_admin/core/network"

//请参考gin 用法
func RegistersWeb() {
	network.Wroutes.LoadHTMLGlob("views/*")
	network.Wroutes.Static("/static", "./static")
	//open websocket
	//network.Wroutes.GET("/socket.io", network.WebsocketIo)
}