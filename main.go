package main

import (
	"ARCTIC-WOLF/setup"
	"ARCTIC-WOLF/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	utils.InitEnv()
	port := utils.EnvData.HttpPort
	if utils.EnvData.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := setup.InitRoutConfig()
	router.Run(":" + strconv.Itoa(port))
}
