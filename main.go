package main

import (
	"TrelloReportTools/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.Routers

	routers.SetupRouters()
	r.Run(":3000")
}
