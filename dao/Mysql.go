package dao

import "github.com/jinzhu/gorm"

var Mysql *gorm.DB

func init() {
	//数据库连接
	var err error
	Mysql, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/article_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}
