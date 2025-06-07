package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "valo-skin-lib-api",
		})
	})
	if err := engine.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		os.Exit(1)
	}
}
