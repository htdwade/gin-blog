package routes

import (
	"gin-blog/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	router := r.Group("/api/v1")
	{
		router.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	_ = r.Run(config.HttpPort)

}