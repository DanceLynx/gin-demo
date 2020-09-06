package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"hello/config"
	"hello/model"
	"strings"
)

var db *gorm.DB

func ConnectDB() {
	database_url := config.DB.DATABASE_URL
	dsn := strings.Split(database_url, "://")
	if len(dsn) < 2 {
		panic("database_url config error:" + database_url)
	}
	var err error
	db, err = gorm.Open(dsn[0], dsn[1])
	if err != nil {
		panic(err)
	}

	//设置与数据库建立连接的最大数目
	db.DB().SetMaxOpenConns(config.DB.MaxOpenConns)
	//设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(config.DB.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DB.ConnMaxLifeTime)

	fmt.Println("connnect to mysql database successful.")
	// 启用Logger，显示详细日志
	db.LogMode(config.DB.ShowSql)
}

func DisconnectDB() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func AutoMigrate() {

	//设置表默认属性
	table_options := "CHARSET=" + config.DB.CHARSET

	db.Set("gorm:table_options", table_options).AutoMigrate(&model.User{})
	fmt.Println("migrate table successful")
}
