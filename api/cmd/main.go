package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tayuchi/valo-skin-lib/api/server"
	"github.com/tayuchi/valo-skin-lib/api/skin"
)

const skinsEndpoint = "https://valorant-api.com/v1/weapons/skins"

func main() {
	engine := gin.Default()
	service := skin.NewSkinService(skinsEndpoint)
	server.SkinRoute(engine, service)

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "valo-skin-lib-api",
		})
	})
	if err := engine.Run(); err != nil { 
		os.Exit(1)
	}
}
