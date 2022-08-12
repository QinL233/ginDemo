package model

import (
	"fmt"
	"ginDemo/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //日志级别
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //取消表明被加s
			},
			DisableForeignKeyConstraintWhenMigrating: true, //取消外键约束
			SkipDefaultTransaction:                   true, //禁用默认事务可以提升性能
		})

	if err != nil {
		log.Fatalf("初始化数据源失败: %v", err)
	}

	g, err := DB.DB()

	if err != nil {
		log.Fatalf("初始化数据源失败: %v", err)
	}
	g.SetMaxIdleConns(10)
	g.SetMaxOpenConns(100)

}
