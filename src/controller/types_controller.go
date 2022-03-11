package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetTypes(c *gin.Context) {
	tc, err := service.GetTypeContent(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "types", gin.H{
		"title": "type",
		"data":  tc,
	})
}
