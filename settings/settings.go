package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 使用viper来管理配置

//  使用内部定义的结构体来存储配置

// Init 这是一个人的项目
func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)
	if err = viper.ReadInConfig(); err != nil {
		// 读取配置信息失败
		fmt.Println("viper.ReadInConfig() failed,err", err)
		return
	}

	// 热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return
}
