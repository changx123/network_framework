package main

import (
	"fmt"
	"network_framework/app/base"
	"os"
	"strconv"
)

func main() {
	t := make(chan int,1)
	base.Run()
	f,err := os.Create("./process.pid")
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte(strconv.Itoa(os.Getpid())))
	}
	f.Close()
	<- t
}