package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayuchi/valo-skin-lib/api/skin"
)

func SkinRoute(r *gin.Engine, service skin.SkinService) {
	r.GET("/skins", func(ctx *gin.Context) {
		data, err := service.GetSkinDataList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get skin data",
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})
}