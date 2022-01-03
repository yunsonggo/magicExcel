package conf

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var AppConf = new(AppConfig)

// LoadConf 加载配置文件
func LoadConf(path ...string) (err error) {
	var configPath string
	if len(path) == 0 {
		flag.StringVar(&configPath, "C", "", "")
		flag.Parse()
		if configPath == "" {
			configPath = "./conf/"
			fmt.Printf("加载默认配置文件:%v\n", configPath)
		} else {
			fmt.Printf("加载命令行指定配置文件:%v\n", configPath)
		}
	} else {
		configPath = path[0]
		fmt.Printf("加载程序内指定的配置文件路径:%v\n", configPath)
	}

	//viper.SetConfigName("config") // 指定配置文件名称 不需要后缀名
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(AppConf); err != nil {
		return
	}
	fmt.Printf("conf:%v\n", AppConf)
	viper.WatchConfig()
	// 配置文件监听回调
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Profile changed:%v\n", e.Name)
		if err = viper.Unmarshal(AppConf); err != nil {
			return
		}
		fmt.Printf("config changed")
	})
	return
}


