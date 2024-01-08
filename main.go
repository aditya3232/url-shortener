package main

import (
	"fmt"

	"github.com/aditya3232/url-shortener/config"
	"github.com/aditya3232/url-shortener/helper"
	"github.com/aditya3232/url-shortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	defer helper.RecoverPanic()

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.Initialize(router)
	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))
}
