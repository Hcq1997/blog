package mapper

import (
	"blog/src/config"
	"errors"
	"fmt"
	"time"
)

type BlogContent struct {
	Id          int
	UserId      int
	Type        int
	TypeName    string `xorm:"-"`
	Tittle       string
	Content     string
	View        int
	IsPublished bool
	Flag        string
	CreateTime  time.Time
	UpdateTime  time.Time
}

type Types struct {
	Type      int
	TypeName  string `xorm:"-"`
	BlogCount int
}

type BlogByYear struct {
	Year        string
	Id          int
	UserId      int
	Type        int
	TypeName    string `xorm:"-"`
	Tittle       string
	Content     string
	View        int
	IsPublished bool
	CreateTime  time.Time
	UpdateTime  time.Time
}

type RecentBlog struct {
	Id       int
	Tittle   string
}

func GetTypeOrderByCountMapper(id int) []Types {
	engine := config.GetDefaultDatabase()
	ts := make([]Types, 0)
	var err error
	if id > 0 {
		err = engine.SQL("SELECT type, count(type) as blog_count "+
			"FROM `blog_content` WHERE user_id = ? GROUP BY type ORDER BY blog_count", id).Find(&ts)
	} else {
		err = engine.SQL("SELECT type, count(type) as blog_count FROM `blog_content` " +
			"GROUP BY type ORDER BY blog_count").Find(&ts)
	}
	if err != nil {
		fmt.Println(err)
	}
	return ts
}

func GetTotalBlogMapper(id int) int {
	engine := config.GetDefaultDatabase()
	bc := new(BlogContent)
	total, err := engine.Table("blog_content").Where("user_id = ?", id).Count(&bc)
	if err != nil {
		fmt.Println(err)
	}
	return int(total)
}

func GetAllBlogMapper(id int) []BlogContent {
	engine := config.GetDefaultDatabase()
	blogs := make([]BlogContent, 0)
	var err error
	if id > 0 {
		err = engine.Table("blog_content").Where("user_id = ?", id).Desc("update_time").
			Find(&blogs)
	} else {
		err = engine.Table("blog_content").Desc("update_time").Find(&blogs)
	}
	if err != nil {
		fmt.Println(err)
	}
	return blogs
}

func GetRecentBlog(id int) []RecentBlog {
	engine := config.GetDefaultDatabase()
	blogs := make([]RecentBlog, 0)
	var err error
	if id > 0 {
		err = engine.Table("blog_content").Where("user_id = ?", id).
			Desc("update_time").Limit(8).Find(&blogs)
	} else {
		err = engine.Table("blog_content").Desc("update_time").Limit(8).Find(&blogs)
	}
	if err != nil {
		fmt.Println(err)
	}
	return blogs
}

func GetBlogByTypeMapper(id, typeId int) []BlogContent {
	engine := config.GetDefaultDatabase()
	blogs := make([]BlogContent, 0)
	var err error
	if id > 0 {
		err = engine.Table("blog_content").Where("user_id = ? and type = ?", id, typeId).
			Desc("update_time").Find(&blogs)
	} else {
		err = engine.Table("blog_content").Where("type = ?", typeId).
			Desc("update_time").Find(&blogs)
	}
	if err != nil {
		fmt.Println(err)
	}
	return blogs
}

func GetBlogArchivesMapper(id int) []BlogByYear {
	engine := config.GetDefaultDatabase()
	blogArchives := make([]BlogByYear, 0)
	var err error
	if id > 0 {
		err = engine.SQL("SELECT DATE_FORMAT(create_time,'%Y') `year`, id, user_id, type, title, "+
			"content, create_time, update_time, `view` FROM blog_content WHERE user_id = ? Order BY "+
			"create_time DESC", id).Find(&blogArchives)
	} else {
		err = engine.SQL("SELECT DATE_FORMAT(create_time,'%Y') `year`, id, user_id, type, title, " +
			"content, create_time, update_time, `view` FROM blog_content Order BY create_time DESC").
			Find(&blogArchives)
	}
	if err != nil {
		fmt.Println(err)
	}
	return blogArchives
}

func InsertBlog(bc *BlogContent) error {
	if bc.UserId <= 0 {
		return errors.New("IllegalUser")
	}
	engine := config.GetDefaultDatabase()
	_, err := engine.Table("blog_content").Insert(bc)
	return err
}

func GetBlogById(userId, blogId int) *BlogContent {
	engine := config.GetDefaultDatabase()
	blog := &BlogContent{}
	engine.Table("blog_content").Where("user_id = ? and id = ? and is_published = 1", userId, blogId).
		Get(blog)
	return blog
}
