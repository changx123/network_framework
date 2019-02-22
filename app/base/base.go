package base

import (
	"blog/core/network"
	"blog/app/routes"
)

func Run()  {
	routes.RegistersWeb()
	routes.RegistersWebSocket()
	routes.RegistersStock()
	network.WRun()
}