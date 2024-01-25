package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var env string
var rootCmd = &cobra.Command{
	Use: "mygo",
}

func main() {
	rootCmd.Flags().StringVarP(&env, "env", "e", "", "环境变量")
	rootCmd.Execute()

	// 获取运行环境 Development  Staging or Production
	environmentName := env //"Development"

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
}

type EmailOptions struct {
	Sender   string
	UserName string
	Password string
}
