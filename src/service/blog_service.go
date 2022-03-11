package service

import (
	"blog/src/mapper"
	"blog/src/util"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Blog struct {
	Username    string
	Visitor     string
	BlogContent *mapper.BlogContent
}

func GetBlogService(c *gin.Context) (*Blog, error) {

	user := c.Param("name")
	userId := mapper.GetIdByUserNameMapper(user)
	blogId := c.Param("id")
	id, err := strconv.Atoi(blogId)
	if err != nil {

	}
	blog := mapper.GetBlogById(userId, id)
	visitor, _ := util.GetUserFromCookie(c)
	ans := &Blog{
		Username:    user,
		Visitor:     visitor,
		BlogContent: blog,
	}
	if blog == nil {
		return nil, errors.New("博客不存在")
	}
	blog.TypeName = mapper.GetTypeNameById(blog.Type)
	return ans, nil
}
