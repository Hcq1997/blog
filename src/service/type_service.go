package service

import (
	"blog/src/mapper"
	"blog/src/util"
	"errors"
	"github.com/gin-gonic/gin"
)

type TypesContent struct {
	UserId        int
	Count         int    // 分类总数
	Username      string // 用户名称
	Visitor       string
	CurrentType   string               // 当前选中的分类
	Types         []mapper.Types       // 分类与博文数量
	SelectedBlogs []mapper.BlogContent // 当前分类博文数组
}

func GetTotalType(id int) int {
	return mapper.GetTotalTypeMapper(id)
}

func GetBlogByType(id int, typeName string) []mapper.BlogContent {
	// 先拿到type id
	typeId := mapper.GetIdByNameMapper(id, typeName)
	// 根据typeId, id 获取博文，时间倒序
	return mapper.GetBlogByTypeMapper(id, typeId)
}

func GetTypeContent(c *gin.Context) (*TypesContent, error) {
	user := c.Param("name")
	// 结构体赋值
	if user == "" {
		return nil, errors.New("用户不存在")
	}
	id := GetIdByPinyin(user)
	if id <= 0 {
		return nil, errors.New("用户不存在")
	}
	// 当前选择的分类，默认为博客数量最多的分类
	// 从前端接受点击的index
	curIndex := 0

	tpContent := &TypesContent{}

	tpContent.Visitor, _ = util.GetUserFromCookie(c)

	tpContent.Username = GetUsernameById(id)
	tpContent.UserId = id
	tpContent.Count = GetTotalType(id)
	tpContent.Types = GetTypeOrderByCount(id)
	if len(tpContent.Types) >= curIndex {
		tpContent.SelectedBlogs = GetBlogByType(id, tpContent.Types[curIndex].TypeName)
	}
	return tpContent, nil
}
