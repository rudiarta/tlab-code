package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/tlab-code/route"
)

func main() {
	app := gin.Default()
	route.InitRoutes(app)
	app.Run()
}
