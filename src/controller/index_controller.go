package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexContent 渲染到模板上的结构体


func GetIndex(c *gin.Context) {

	indexContent, err := service.GetIndexContent(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "index",
		"data":  indexContent,
	})
}
