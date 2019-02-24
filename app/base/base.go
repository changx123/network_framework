package base

import (
	"blog_admin/core/network"
	"blog_admin/app/routes"
)

func Run()  {
	routes.RegistersWeb()
	routes.RegistersWebSocket()
	routes.RegistersStock()
	network.WRun()
}