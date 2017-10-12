package tables

import (
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
	"github.com/spf13/viper"
)

func GetDBFromVipper() *gorm.DB {
	log.Print("开始初始化数据库连接")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetString("db_name"))
	DB, err := gorm.Open("mysql", dsn+"&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("数据库连接错误:", err)
		return nil
	}

	log.Print("数据库连接初始化成功")

	return DB
}
