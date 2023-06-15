package model

import (
	"AM_system/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error
var dialector gorm.Dialector

func InitDb() {
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//此处以Mysql为例
	dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	db, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	sqlDB, err := db.DB()
	//defer，golang的语法糖，在最后执行，值得一提的是，gorm.db没有写close方法来关闭数据库，只能使用db.DB()获取到底层的sql类型再关闭数据库
	//关闭数据库
	//defer sqlDB.Close()

	//自动迁移，将结构体中的表格迁移到Mysql中，但在这个项目中我们不需要这个，因为我是已经在Mysql中建过表格了的
	//db.AutoMigrate()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	if err != nil {
		fmt.Println("返回数据库接口失败，请检查参数：", err)
	}
}
