package main

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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

	// 从appsettings.json读取
	name := viper.GetString("AppName")
	fmt.Println(name)

	// 从appsettings.{environmentName}.json读取
	mysqlConnectionString := viper.GetString("ConnectionStrings.Mysql")
	fmt.Println(mysqlConnectionString)

	// Options模式使用struct提供对相关配置组的强类型访问
	var emailOptions EmailOptions
	viper.UnmarshalKey("EmailOptions", &emailOptions)
	fmt.Println(emailOptions.Sender)
	fmt.Println(emailOptions.UserName)
	fmt.Println(emailOptions.Password)

	// 读取远程consul配置
	consulEndPoint := viper.GetString("ConsulEndPoint")
	viper.AddRemoteProvider("consul", consulEndPoint, "section1.json")
	viper.ReadRemoteConfig()

	// 从远程consul读取
	fmt.Println(viper.GetString("Level"))
}

type EmailOptions struct {
	Sender   string
	UserName string
	Password string
}
