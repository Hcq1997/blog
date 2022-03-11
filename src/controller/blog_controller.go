package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBlog(c *gin.Context)  {
	blog, err := service.GetBlogService(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "blog-detail", gin.H{
		"title": "博客详情",
		"data":  blog,
	})
}

