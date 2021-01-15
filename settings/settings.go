package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 配置viper读取的参数信息，并读取配置文件
func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
	})
	return
}
