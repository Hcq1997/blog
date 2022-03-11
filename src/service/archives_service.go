package service

import (
	"blog/src/mapper"
	"blog/src/util"
	"errors"
	"github.com/gin-gonic/gin"
)

type BlogArchives struct {
	Year  string
	Blogs []mapper.BlogContent
}

type ArchivesContent struct {
	Id        int
	Username  string
	Visitor   string
	TotalBlog int
	Blogs     []BlogArchives
}

// GetBlogArchives 按照时间归档博客
func GetBlogArchives(id int) []BlogArchives {
	blogByYear := make([]mapper.BlogByYear, 0)
	blogByYear = mapper.GetBlogArchivesMapper(id)

	ba := make([]BlogArchives, 0)
	// 按年归档，填入类型
	idx := 0
	for _, blog := range blogByYear {
		blogContent := mapper.BlogContent{
			Id:         blog.Id,
			UserId:     blog.UserId,
			Type:       blog.Type,
			TypeName:   mapper.GetTypeNameById(blog.Type),
			Tittle:     blog.Tittle,
			Content:    blog.Content,
			View:       blog.View,
			CreateTime: blog.CreateTime,
			UpdateTime: blog.UpdateTime,
		}
		// 如果为该年的第一条博客，先创建
		if len(ba) <= idx || ba[idx-1].Year != blog.Year {
			ba = append(ba, BlogArchives{
				Year:  blog.Year,
				Blogs: make([]mapper.BlogContent, 0),
			})
			idx++
		}
		// 将内容放进该年度
		ba[idx-1].Blogs = append(ba[idx-1].Blogs, blogContent)
	}

	return ba
}

func GetArchivesContent(c *gin.Context) (*ArchivesContent, error) {
	user := c.Param("name")
	// 结构体赋值
	if user == "" {
		return nil, errors.New("用户不存在")
	}
	id := GetIdByPinyin(user)
	if id <= 0 {
		return nil, errors.New("用户不存在")
	}

	archivesContent := &ArchivesContent{}

	archivesContent.Visitor, _ = util.GetUserFromCookie(c)

	archivesContent.Id = id
	archivesContent.Username = GetUsernameById(id)
	archivesContent.TotalBlog = GetTotalBlog(id)
	archivesContent.Blogs = GetBlogArchives(id)
	return archivesContent, nil
}
