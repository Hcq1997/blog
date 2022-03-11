package service

import (
	"blog/src/config"
	"blog/src/mapper"
	"blog/src/util"
	"github.com/gin-gonic/gin"
	"html/template"
)

type LogInfo struct {
	Succeed bool
	Username    string
	Message string
	Access  template.HTML // 用来控制html标签显示隐藏，只有hidden与""两个值
}

func LogPage(c *gin.Context) *LogInfo {
	logInfo := &LogInfo{
		Access: "Hidden",
	}
	return logInfo
}

func Login(c *gin.Context) *LogInfo {
	var user, password string
	logInfo := &LogInfo{}
	if username, has := c.GetPostForm("username"); has {
		user = username
	} else {
		logInfo.Message = "未获取到您的用户名"
	}
	if p, has := c.GetPostForm("password"); has {
		password = p
	} else {
		logInfo.Message = "未获取到您的密码"
	}
	if util.AuthorityPassword(user, password) {
		// 校验通过
		// 跳转页面
		logInfo.Succeed = true
		logInfo.Username = user
		c.SetCookie("user", user, config.CookieAge, config.CookiePath, config.CookieDomain,
			config.CookieSecure, config.CookieHttpOnly)
		c.SetCookie("password", password, config.CookieAge, config.CookiePath, config.CookieDomain,
			config.CookieSecure, config.CookieHttpOnly)

	} else {
		logInfo.Succeed = false
		logInfo.Message = "用户名或密码错误，如有需要，请联系QQ:1695524096"
		logInfo.Access = ""
	}
	// 记录日志
	return logInfo
}

func GetIndexContentAfterLogin(username string) *IndexContent {
	indexContent := &IndexContent{}
	user := username
	id := mapper.GetIdByUserNameMapper(user)
	indexContent.Id = id
	// 此处不需要校验是否登录
	indexContent.Visitor = user
	indexContent.TotalBlog = GetTotalBlog(id)
	indexContent.Username = username
	indexContent.Blogs = GetAllBlogById(id)
	indexContent.RecentBlog = GetRecentBlog(id)
	indexContent.TypeCount = GetTypeOrderByCount(id)
	return indexContent
}
