package base

import (
<<<<<<< HEAD
	"network_framework/core/network"
	"network_framework/app/routes"
=======
	"blog_admin/core/network"
	"blog_admin/app/routes"
>>>>>>> c2c3f0ec19d9d0aa9469733956311c89d426dea1
)

func Run()  {
	routes.RegistersWeb()
	routes.RegistersWebSocket()
	routes.RegistersStock()
	network.WRun()
}