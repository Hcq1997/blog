package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)



func GetArchives(c *gin.Context) {

	archivesContent, err := service.GetArchivesContent(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "archives", gin.H{
		"title": "archives",
		"data": archivesContent,
	})
}
