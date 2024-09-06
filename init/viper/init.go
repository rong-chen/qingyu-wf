package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"qinyu-wf/init/db"
)

func Init() (db db.Config) {
	env := os.Getenv("APP_ENV") // 读取环境变量
	if env == "" {
		env = "development" // 默认值
	}
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	switch env {
	case "production":
		v.SetConfigName("config.pro")
	case "development":
		v.SetConfigName("config.dev")
	default:
		panic("未知环境，请设置成开发或者生产环境")
	}
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取配置失败，请重新设置配置信息:%s`, err))
		return
	}
	if err := v.Unmarshal(&db); err != nil {
		panic(err)
	}
	return
}
