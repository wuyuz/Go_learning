package common

import (
	"data_worker/app"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

var ViperConfig *viper.Viper //nolint:gochecknoglobals

func LoadEnv() {
	ViperConfig = viper.New()
	ViperConfig.AddConfigPath("./")   // 第一个搜索路径，可以设置多个AddConfigPath
	ViperConfig.SetConfigFile(".env") // 设置配置文件，不太后缀
	ViperConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_ = ViperConfig.ReadInConfig() // 搜索路径，并读取配置数据
	ViperConfig.AutomaticEnv()     // 表示先预加载匹配本机的环境变量，可以判断本地环境变量，开发环境
	ViperConfig.WatchConfig()      // Viper支持让你的应用程序在运行时拥有读取配置文件的能力
	viper.OnConfigChange(func(e fsnotify.Event) { // 监控配置变动
		app.Log.Info().Msg("Config file changed")
	})
}
