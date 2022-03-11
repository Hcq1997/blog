package mapper

import (
	"blog/src/config"
	"fmt"
	"testing"
)

func TestGetTotalTypeMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	ans := GetTotalTypeMapper(7)
	fmt.Println(ans)
}

func TestGetIdByNameMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	ans := GetIdByNameMapper(0,"fds")
	fmt.Println(ans)
}

func TestGetAllTypeMapper(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	fmt.Println(GetAllTypeMapper(1))
}

func TestGetTypeNameById(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	name := GetTypeNameById(3)
	fmt.Println(name)
}