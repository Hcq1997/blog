package service

import (
	"blog/src/mapper"
	"blog/src/util"
	"errors"
	"github.com/gin-gonic/gin"
)

type IndexContent struct {
	Id         int
	TotalBlog  int
	Username   string
	Visitor    string
	Blogs      []mapper.BlogContent
	RecentBlog []mapper.RecentBlog
	TypeCount  []mapper.Types
}

// GetTotalBlog 获取博客总数
func GetTotalBlog(id int) int {
	total := mapper.GetTotalBlogMapper(id)
	return total
}

func GetAllBlogById(id int) []mapper.BlogContent {
	return mapper.GetAllBlogMapper(id)
}

func GetRecentBlog(id int) []mapper.RecentBlog {
	return mapper.GetRecentBlog(id)
}

func GetTypeOrderByCount(id int) []mapper.Types {
	ans := make([]mapper.Types, 0)
	ans = mapper.GetTypeOrderByCountMapper(id)
	for i := 0; i < len(ans); i++ {
		ans[i].TypeName = mapper.GetTypeNameById(ans[i].Type)
	}
	return ans
}

func GetUsernameById(id int) string {
	return mapper.GetUsernameByIdMapper(id)
}

func GetIdByPinyin(user string) int {
	return mapper.GetIdByUserNameMapper(user)
}

func GetIndexContent(c *gin.Context) (*IndexContent, error) {
	indexContent := &IndexContent{}
	user := c.Param("name")
	// 结构体赋值
	if user == "" {
		return nil, errors.New("用户不存在")
	}
	id := GetIdByPinyin(user)
	if id <= 0 {
		return nil, errors.New("用户不存在")
	}
	indexContent.Id = id
	// 此处不需要校验是否登录
	indexContent.Visitor, _ = util.GetUserFromCookie(c)
	indexContent.TotalBlog = GetTotalBlog(id)
	indexContent.Username = GetUsernameById(id)
	indexContent.Blogs = GetAllBlogById(id)
	indexContent.RecentBlog = GetRecentBlog(id)
	indexContent.TypeCount = GetTypeOrderByCount(id)
	return indexContent, nil
}
