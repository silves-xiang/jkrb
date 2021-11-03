package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	MysqlAdd      string `yaml:"MysqlAdd"`
	MysqlPort     int `yaml:"MysqlPort"`
	MysqlDatabase string `yaml:"MysqlDatabase"`
	MysqlUser string `yaml:"MysqlUser"`
	MysqlPwd string `yaml:"MysqlPwd"`
	EmailPoster string `yaml:"EmailPoster"`
	EmailPwd string `yaml:"EmailPwd"`
	Time string `yaml:"Time"`
	Minites string `yaml:"Minites"`
}

var ConfigA Config
func CreateConfig () Config {
	file , err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Println(err)
	}
	var ConfigR Config
	err = yaml.Unmarshal(file, &ConfigR)
	if err != nil {
		log.Println(err)
	}
	ConfigA = ConfigR
	fmt.Println(ConfigR)
	return ConfigR
}