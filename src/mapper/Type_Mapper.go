package mapper

import (
	"blog/src/config"
	"fmt"
	"time"
)

type BlogType struct {
	UserId     int
	Type       string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetTotalTypeMapper(id int) int {
	engine := config.GetDefaultDatabase()
	var count int64
	var err error
	if id > 0 {
		count, err = engine.Table("blog_type").Where("user_id = ?", id).Count("id")
	} else {
		count, err = engine.Table("blog_type").Count()
	}
	if err != nil {
		fmt.Println(err)
	}
	return int(count)
}

func GetIdByNameMapper(id int, typeName string) int {
	engine := config.GetDefaultDatabase()
	var typeId int
	var err error
	var ok bool
	if id > 0 {
		ok, err = engine.Table("blog_type").Cols("id").
			Where("user_id = ? and type = ?", id, typeName).Get(&typeId)
	} else {
		ok, err = engine.Table("blog_type").Cols("id").Where("type = ?", typeName).
			Get(&typeId)
	}
	if !ok {
		fmt.Println(err)
	}
	return typeId
}

// GetTypeNameById 根据类型id获取类型名
func GetTypeNameById(id int) string {
	engine := config.GetDefaultDatabase()
	var name string
	var err error
	var ok bool
	if id > 0 {
		ok, err = engine.Table("blog_type").Cols("type").Where("id = ?", id).Get(&name)
	} else {
		ok, err = engine.Table("blog_type").Cols("type").Get(&name)
	}
	if !ok {
		fmt.Println(err)
	}
	return name
}

// GetAllTypeMapper 拿到所有分类
func GetAllTypeMapper(id int) []string {
	engine := config.GetDefaultDatabase()
	types := make([]string, 0)
	if id > 0 {
		engine.Table("blog_type").Cols("type").Where("user_id = ?", id).Find(&types)
	} else {
		engine.Table("blog_type").Cols("type").Find(&types)
	}
	return types
}

// CheckType 检查分类是否存在
func CheckType(id int, typeName string) (bool, int) {
	var userId int
	var has bool
	engine := config.GetDefaultDatabase()
	if id > 0 {
		has, _ = engine.Table("blog_type").Cols("id").
			Where("user_id = ? and type = ?", id, typeName).Get(&userId)
	} else {
		has, _ = engine.Table("blog_type").Cols("id").
			Where(" type = ?", typeName).Get(&userId)
	}
	return has, userId
}

// InsertNewType 新建分类
func InsertNewType(blogType *BlogType) (bool, int) {
	if blogType.UserId < 0 {
		fmt.Println("id < 0")
		return false, 0
	}
	engine := config.GetDefaultDatabase()

	var err error
	_, err = engine.Table("blog_type").Insert(blogType)
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	typeId := GetIdByNameMapper(blogType.UserId, blogType.Type)
	return true, typeId
}

