package service

import (
	"blog/src/mapper"
	"blog/src/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type ManagementList struct {
	Tittle     string
	Type       int
	TypeName   string
	Published  string
	UpdateTime time.Time
}

type Management struct {
	Types    []string
	Username string
	ManaList []ManagementList
}

type BlogInput struct {
	Id       int
	Username string
	Types    []string
}

type PreviewBlog struct {
	Username string
	Blog     *mapper.BlogContent
}

// GetAllType 拿到所有分类
func GetAllType(id int) []string {
	return mapper.GetAllTypeMapper(id)
}

func SaveBlog(c *gin.Context) *mapper.BlogContent {
	blog := &mapper.BlogContent{}
	blog.UserId = 1
	if flag, has := c.GetPostForm("flag"); has {
		blog.Flag = flag
	} else {
		blog.Flag = "原创"
	}
	// 前端做校验，一般不会出现未输入博客标题或者内容的问题
	if tittle, has := c.GetPostForm("title"); has {
		blog.Tittle = tittle
	} else {
		blog.Tittle = fmt.Sprintf("%s%s%d", mapper.GetUsernameByIdMapper(blog.Id), "的博客", time.Now().Unix())
	}
	if content, has := c.GetPostForm("content"); has {
		blog.Content = content
	} else {
		blog.Content = "该博客不存在"
	}
	if blogType, has := c.GetPostForm("type"); has {
		blog.TypeName = blogType
	} else {
		blog.TypeName = "未分类"
	}
	// 如果类型不存在，则新建类型
	if has, typeId := mapper.CheckType(blog.UserId, blog.TypeName); has {
		blog.Type = typeId
	} else {
		t := &mapper.BlogType{
			UserId:     blog.UserId,
			Type:       blog.TypeName,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		_, ti := mapper.InsertNewType(t)
		blog.Type = ti
	}
	blog.Type = mapper.GetIdByNameMapper(blog.UserId, blog.TypeName)
	if published, has := c.GetPostForm("published"); has && published == "true" {
		blog.IsPublished = true
	} else {
		blog.IsPublished = false
	}
	blog.CreateTime = time.Now()
	blog.UpdateTime = time.Now()
	// 持久化
	mapper.InsertBlog(blog)
	return blog
}

func GetBlogManagementList(c *gin.Context) (*Management, error) {
	au := util.AuthorityUser(c)
	if !au {
		return nil, errors.New("用户未登录")
	}
	ml := make([]ManagementList, 0)
	id := 1
	blog := mapper.GetAllBlogMapper(id)
	for _, val := range blog {
		item := ManagementList{
			Tittle:     val.Tittle,
			Type:       val.Type,
			TypeName:   mapper.GetTypeNameById(val.Type),
			UpdateTime: val.UpdateTime,
		}
		if val.IsPublished {
			item.Published = "是"
		} else {
			item.Published = "否"
		}
		ml = append(ml, item)
	}
	name, _ := util.GetUserFromCookie(c)
	mana := &Management{
		ManaList: ml,
		Username: name,
		Types:    mapper.GetAllTypeMapper(id),
	}
	return mana, nil
}

func SearchBlogsService(c *gin.Context) Management {
	// 获得筛选条件 title type
	blogType, hasType := c.GetPostForm("type")
	blogTitle, hasTitle := c.GetPostForm("title")
	fmt.Println(blogType, blogTitle)
	ml := make([]ManagementList, 0)
	name, _ := util.GetUserFromCookie(c)
	id := mapper.GetIdByUserNameMapper(name)
	var blog []mapper.BlogContent
	if hasType && blogType != "分类" {
		typeId := mapper.GetIdByNameMapper(id, blogType)
		blog = mapper.GetBlogByTypeMapper(id, typeId)
	} else {
		blog = mapper.GetAllBlogMapper(id)
	}
	for _, val := range blog {
		// 筛选出符合条件的博客
		// 不作筛选
		if !hasTitle || strings.Contains(val.Tittle, blogTitle) {
			item := ManagementList{
				Tittle:     val.Tittle,
				Type:       val.Type,
				TypeName:   mapper.GetTypeNameById(val.Type),
				UpdateTime: val.UpdateTime,
			}
			if val.IsPublished {
				item.Published = "是"
			} else {
				item.Published = "否"
			}
			ml = append(ml, item)
		}
	}
	mana := Management{
		ManaList: ml,
		Username: name,
		Types:    mapper.GetAllTypeMapper(id),
	}
	return mana
}

func GetBlogInput(c *gin.Context) (*BlogInput, error) {
	// 需要先校验用户
	au := util.AuthorityUser(c)
	if !au {
		return nil, errors.New("用户未登录")
	}
	user, _ := util.GetUserFromCookie(c)

	blogInput := &BlogInput{}
	blogInput.Id = mapper.GetIdByUserNameMapper(user)
	blogInput.Username = user
	blogInput.Types = GetAllType(blogInput.Id)
	return blogInput, nil
}

func Preview(c *gin.Context) *PreviewBlog {
	blog := SaveBlog(c)
	pb := &PreviewBlog{
		Username: c.Param("name"),
		Blog:     blog,
	}
	return pb
}
