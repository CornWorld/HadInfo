package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Bootstrap() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.StaticFS("/static", http.Dir("static"))

	r.LoadHTMLGlob("page/*")

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/add", handleAddForm)
	r.GET("/api/v/:pasteId", handleViewPasteForm)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", gin.H{})
	})
	r.GET("/v/:pasteId", handleViewPastePage)

	err := r.Run(":3000") // TODO: add configure options
	if err != nil {
		return
	}
}
