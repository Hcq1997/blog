package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	GetConfig("config.yaml")
	fmt.Println(Conf)
}
