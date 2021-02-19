package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort         int `json:"apiPort"`
	ApiReadTimeout  int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
	EtcdEndPoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
	Web string `json:"webroot"`
	MongodbUri string `json:"mongodbUri"`
	MongodbConnectTimeout int `json:"mongodbConnectTimeout"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 读取配置文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// json反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	// 赋值单例
	G_config = &conf
	return
}
