package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 一个数据库的连接
func init() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
