package controller

import (
	"blog/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	logIn := service.LogPage(c)
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "登录博客",
		"data": logIn,
	})
}

func LoginSuccess(c *gin.Context) {
	logInfo := service.Login(c)

	if !logInfo.Succeed {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "登陆失败",
			"data": logInfo,
		})
		return
	}
	indexContent := service.GetIndexContentAfterLogin(logInfo.Username)
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "主页",
		"data":  indexContent,
	})
}
