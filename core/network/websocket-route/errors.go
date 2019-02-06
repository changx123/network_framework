package websocket_route

import "errors"

var (
	ERROR_STOP = errors.New("websocket-route: STOP")

	ERROR_CONTINUE = errors.New("websocket-route: CONTINUE")

	ERROR_BREAK = errors.New("websocket-route: BREAK")
)
