package mapper

import (
	"blog/src/config"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"time"
)

func GetUsernameByIdMapper(id int) string {
	engine := config.GetDefaultDatabase()
	var username string
	var ok bool
	var err error
	if id > 0 {
		ok, err = engine.Table("blog_user").Cols("nickname").Where("id = ?", id).Get(&username)
	} else {
		uid, uError := uuid.GenerateUUID()
		if uError != nil {
			uid = fmt.Sprintf("%d", time.Now().Unix())
		}
		username = fmt.Sprintf("%s%s", "访客", uid)
	}
	if !ok {
		fmt.Println(err)
	}
	return username
}

func GetIdByUserNameMapper(user string) int {
	engine := config.GetDefaultDatabase()
	var id int
	engine.Table("blog_user").Cols("id").Where("pin_yin = ?", user).Get(&id)
	return id
}

func GetPassword(user string) string {
	var p string
	engine := config.GetDefaultDatabase()
	engine.Table("blog_user").Cols("password").Where("nickname = ?", user).Get(&p)
	return p
}
