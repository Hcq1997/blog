package mapper

import (
	"blog/src/config"
	"fmt"
	"testing"
)

func TestGetTotalBlogMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	fmt.Println(GetTotalBlogMapper(-1))
}

func TestGetAllBlogMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	ans := GetAllBlogMapper(1)
	fmt.Println(ans)
}

func TestGetTypeOrderByCountMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	ans := GetTypeOrderByCountMapper(1)
	fmt.Println(ans)
}

func TestGetBlogArchivesMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	ans := GetBlogArchivesMapper(-1)
	fmt.Println(ans)
}
func TestGetBlogById(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	fmt.Println(GetBlogById(1, 0))
}

func TestGetBlogByTypeMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	fmt.Println(GetBlogByTypeMapper(1,1))
}

func TestGetRecentBlog(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	b := GetRecentBlog(1)
	fmt.Println(b)
}
