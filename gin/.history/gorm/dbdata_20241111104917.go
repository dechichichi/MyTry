package dbdata

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

// 一个数据库的连接
func Init() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//创建表  自动迁移
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Tom", Sex: true, Age: 20})
	defer db.Close()
}
