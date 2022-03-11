package util

import (
	"blog/src/mapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func generateUsername() string {
	return fmt.Sprintf("%s%d", "访客", time.Now().Unix())
}

// GetUserFromCookie 通过cookie校验用户
// 返回值 user 用户名  bool 用户是否登录
func GetUserFromCookie(c *gin.Context) (string, bool) {
	user, err := c.Cookie("user")
	var has bool
	if err != nil || user == "" {
		user = generateUsername()
		has = false
	}
	has = true
	return user, has
}

// AuthorityUser 根据cookie校验用户是否登录
func AuthorityUser(c *gin.Context) bool {
	user, errUser := c.Cookie("user")
	password, errPass := c.Cookie("password")
	if errUser != nil || errPass != nil {
		return false
	}
	if mapper.GetPassword(user) != ToMd5(password) {
		return false
	}
	return true
}

// AuthorityPassword 根据密码校验用户
func AuthorityPassword(username, password string) bool {
	if mapper.GetPassword(username) != ToMd5(password) {
		return false
	}
	return true
}
