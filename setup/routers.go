package setup

import (
	"ARCTIC-WOLF/code/handlers"

	"github.com/gin-gonic/gin"
)

func rootRouters(route *gin.Engine) {

	// Ping test
	route.GET("/", pingHandler)

	route.GET("/v1/risks", organizeController(handlers.GetRisks))
	route.POST("/v1/risks", organizeController(handlers.CreateRisk))
	route.GET("/v1/risks/:id", organizeControllerwithGin(handlers.GetRiskById))

}

func pingHandler(c *gin.Context) {
	c.String(200, "pong")
}
