package routes

func RegistersWebSocket() {
	////****中间件
	//network.Wsroutes.Use()
	////****注册状态钩子
	////新连接通知
	//network.Wsroutes.Hook(websocket_route.HOOK_NEW_CONN, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////连接closed 通知
	//network.Wsroutes.Hook(websocket_route.HOOK_CLOSED, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////路由寻址不存在
	//network.Wsroutes.Hook(websocket_route.HOOK_NOT_MODULE, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////错误通知
	//network.Wsroutes.Hook(websocket_route.HOOK_ERROR, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////发送消息通知
	//network.Wsroutes.Hook(websocket_route.HOOK_WRITE_MESSAGE, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////接收消息解包
	//network.Wsroutes.Hook(websocket_route.HOOK_UN_PACKING, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
	////发送消息封包
	//network.Wsroutes.Hook(websocket_route.HOOK_PACKET, func(conn *websocket.Conn, msg *websocket_route.Message, route *interface{}) error {
	//	return nil
	//})
}
