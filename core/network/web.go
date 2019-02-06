package network

import (
	"github.com/gin-gonic/gin"
	"network_framework/config"
)

var Wroutes *gin.Engine

func init() {
	Wroutes = gin.Default()
}

func WRun() {
	if config.HTTPS_OPEN {
		go Wroutes.RunTLS(config.HTTPS_LISTEN_ADDRESS,config.HTTPS_CERTFILE_PATH,config.HTTPS_KEYFILE_PATH)
	}
	if config.HTTP_OPEN {
		go Wroutes.Run(config.HTTP_LISTEN_ADDRESS)
	}
}

