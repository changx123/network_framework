package main

import "network_framework/app/base"

func main() {
	t := make(chan int,1)
	base.Run()
	<- t
}