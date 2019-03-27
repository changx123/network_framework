package base

import (
	"network_framework/core/network"
	"network_framework/app/routes"
)

func Run()  {
	routes.RegistersWeb()
	routes.RegistersWebSocket()
	routes.RegistersStock()
	network.WRun()
}