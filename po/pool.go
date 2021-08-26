package po

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewConnPool() *gorm.DB {
	//配置MySQL连接参数
	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.2.204" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "xitie"       //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("开启连接池失败, error=" + err.Error())
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second * 300)
	return db
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		db = NewConnPool()
		sqlDB, _ = db.DB()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		db = NewConnPool()
	}
	return db
}
