package main

import (
	"blog/src/config"
	"blog/src/controller"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.GetConfig("./src/config/config.yaml")

	// 使用Gin框架
	r := gin.Default()
	r.Static("/res", "src/static/")
	// 使用多模板
	mul := multitemplate.NewRenderer()
	mul.AddFromFiles("index", "src/view/base.tmpl", "src/view/index.tmpl")
	mul.AddFromFiles("types", "src/view/base.tmpl", "src/view/types.tmpl")
	mul.AddFromFiles("archives", "src/view/base.tmpl", "src/view/archives.tmpl")
	mul.AddFromFiles("management", "src/view/base.tmpl", "src/view/blogs-input.tmpl")
	mul.AddFromFiles("management-list", "src/view/base.tmpl", "src/view/blogs.tmpl")
	mul.AddFromFiles("blogs-search", "src/view/base.tmpl", "src/view/blogs-search.tmpl")
	mul.AddFromFiles("login", "src/view/base.tmpl", "src/view/login.tmpl")
	mul.AddFromFiles("blog-detail", "src/view/base.tmpl", "src/view/blog-detail.tmpl")
	mul.AddFromFiles("error", "src/view/base.tmpl", "src/view/error.tmpl")
	mul.AddFromFiles("preview", "src/view/base.tmpl", "src/view/preview.tmpl")
	r.HTMLRender = mul

	// 定义路由
	r.GET("/", controller.FirstPage)

	r.GET("/:name/index", controller.GetIndex)
	r.GET("/:name/type", controller.GetTypes)
	r.GET("/:name/archives", controller.GetArchives)
	r.GET("/:name/management", controller.GetBlogInput)
	r.GET("/:name/management/list", controller.GetBlogList)

	// 搜索博客内容
	r.POST("/blogs-search", controller.SearchBlogs)

	// 登录页面
	r.GET("/login", controller.Login)
	// 点击登陆按钮
	r.POST("/login", controller.LoginSuccess)

	r.GET("/:name/blog/:id", controller.GetBlog)

	r.POST("/:name/management", controller.SaveBlog)

	// 启动服务
	r.Run(":8000")
}
