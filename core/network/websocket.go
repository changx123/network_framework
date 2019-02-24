package network

import (
	"blog_admin/core/network/websocket-route"
	"github.com/gin-gonic/gin"
	"fmt"
)


var Wsroutes *websocket_route.StorageGroup

func init() {
	Wsroutes = websocket_route.NewRouter()
}

func WebsocketIo(context *gin.Context) {
	conn, err := websocket_route.NewConn(context.Writer, context.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	Wsroutes.Listen(conn)
}