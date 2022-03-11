package service

import (
	"blog/src/config"
	"fmt"
	"testing"
)

func TestGetTypeOrderByCount(t *testing.T) {
	config.GetConfig("../config/config.yaml")
	fmt.Println(config.Conf)
	fmt.Println(GetTypeOrderByCount(1))
}
