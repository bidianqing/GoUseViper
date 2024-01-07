package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 获取运行环境 Development  Staging or Production
	environmentName := "Development"

	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	// 公共的配置文件，无论哪个环境都加载
	viper.SetConfigName("appsettings")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 读取对应环境的配置文件
	viper.SetConfigName("appsettings." + environmentName)
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	name := viper.Get("AppName")
	fmt.Println(name)

	mysqlConnectionString := viper.GetString("ConnectionStrings.Mysql")
	fmt.Println(mysqlConnectionString)
}
