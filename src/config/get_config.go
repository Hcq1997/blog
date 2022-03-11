package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Table    string
	}
}

var Conf Config

func GetConfig(filepath string) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err)
	}
}
