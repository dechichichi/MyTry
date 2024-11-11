package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}
