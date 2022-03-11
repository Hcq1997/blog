package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)



// GetBlogInput 输入博客
func GetBlogInput(c *gin.Context) {
	bi, err := service.GetBlogInput(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "management", gin.H{
		"title": "input",
		"data":  bi,
	})
}

func GetBlogList(c *gin.Context) {
	mana, err := service.GetBlogManagementList(c)
	if err != nil {
		c.HTML(http.StatusOK, "error", gin.H{
			"data": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "management-list", gin.H{
		"title": "management-list",
		"data":  mana,
	})
}

func SaveBlog(c *gin.Context) {
	blog := service.Preview(c)
	c.HTML(http.StatusOK, "preview", gin.H {
		"title": "预览页面",
		"data":  blog,
	})
}

// SearchBlogs 博客管理页面搜索博客内容
func SearchBlogs(c *gin.Context) {
	cons := service.SearchBlogsService(c)
	c.HTML(http.StatusOK, "management-list", gin.H{
		"title": "management-list",
		"data":  cons,
	})
}
