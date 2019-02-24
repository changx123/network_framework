package main

import (
	"blog_admin/app/base"
	"blog_admin/config"
	"blog_admin/core/network"
	"runtime"
)

func init() {
	network.UpdatePidFile()
}

func main() {
	if config.MAX_CPUS > 0 {
		runtime.GOMAXPROCS(config.MAX_CPUS)
	}else{
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	t := make(chan int,1)
	base.Run()
	if !config.WEB_DEBUG && config.HTTP_HOT_UPDATE {
		network.SingalHandler()
	}
	<- t
}