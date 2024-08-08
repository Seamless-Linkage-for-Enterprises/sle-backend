package config

import (
	"sle/routes"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Configuration() {
	routes.SetupRoutes(r)
}

func RunServer(listenAddr string) {
	r.Run(listenAddr)
}
