package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"qingyu-wf/api/chat"
	"qingyu-wf/api/classify"
	"qingyu-wf/api/friendRelationship"
	"qingyu-wf/api/user"
	"qingyu-wf/global"
)

type Config struct {
	MySqlConfig MySqlConfig `mapstructure:"MySql"`
}

type MySqlConfig struct {
	User     string `mapstructure:"User"`
	Pwd      string `mapstructure:"Pwd"`
	Database string `mapstructure:"Database"`
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
}

type MigrateInterface interface {
	AutoMigrateFunc()
}

var migrateList = []interface{}{
	&user.User{},
	&friendRelationship.FriendRelationship{},
	&friendRelationship.AwaitingAgreeTable{},
	&classify.TableClassify{},
	&chat.ContentChatTable{},
}

func Init(config Config) {
	mysqlConfig := config.MySqlConfig
	//链接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Pwd, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	global.MySql = db
}

func InitAutoMigrate() {
	for i := range migrateList {
		err := global.MySql.AutoMigrate(migrateList[i])
		if err != nil {
			panic(err)
		}

	}
}
