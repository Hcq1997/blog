package controller

import (
	"blog/src/service"
	"blog/src/util"
	"github.com/gin-gonic/gin"
	template2 "html/template"
	"net/http"
)

func FirstPage(c *gin.Context) {
	isLog := util.AuthorityUser(c)
	user, err := c.Cookie("user")
	if err != nil {

	}
	if !isLog {
		c.HTML(http.StatusOK, "error", gin.H{
			"title": "错误",
			"data": template2.HTML(`<div>您还未登录</div><br><a href="/login">登陆</a>`),
		})
		return
	}
	indexContent := service.GetIndexContentAfterLogin(user)
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "主页",
		"data": indexContent,
	})
}
