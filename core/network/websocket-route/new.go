package websocket_route

import (
	"github.com/changx123/websocket-sync"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     SetWebSocketOrigin,
}

func SetWebSocketOrigin(_ *http.Request) bool {
	return true
}

func NewConn(w http.ResponseWriter,r *http.Request) (*websocket.Conn , error) {
	upgrader.Init()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil , err
	}
	return conn , nil
}


