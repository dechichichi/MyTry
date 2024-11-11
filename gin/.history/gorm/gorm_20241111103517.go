package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
}
