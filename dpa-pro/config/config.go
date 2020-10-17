package utils

import (
	"dpa-pro/lib"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	LDAP_HOST      string `yaml: "LDAP_HOST"`
	LDAP_PORT      int    `yaml: "LDAP_PORT"`
	LDAP_USER_NAME string `yaml: "LDAP_USER_NAME"`
	LDAP_PASSWORD  string `yaml: "LDAP_PASSWORD"`
}

//读取Yaml配置文件, 并转换成conf对象
func (c *config) loadConf() *config {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("config.yaml")
	lib.CheckErr(err, "the config.yaml load failed")

	err = yaml.Unmarshal(yamlFile, c)
	lib.CheckErr(err, "the config data Unmarshal failed")
	return c
}

// 根据键获取配置值
func GetConf(param string) interface{} {
	c := new(config)
	//读取yaml配置文件
	conf := c.loadConf()

	v, err := GetStructStringField(conf,param)
	if err != nil {
		lib.CheckErr(err, "the config data get string value wrong")
	}
	fmt.Println(v,"sdfasfas")

	return v
}
