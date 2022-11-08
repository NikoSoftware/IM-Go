package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func InitViper() {

	viper.SetConfigName("config/app")
	viper.AddConfigPath("./") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	log.Printf("app 配置 %v", viper.Get("app"))

}
