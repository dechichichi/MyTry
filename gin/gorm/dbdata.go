package dbdata

import (
	"fmt"

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
	db.Create(&User{Name: "Tim", Sex: false, Age: 23})
	//查
	var user User
	db.Where("age<?", "20").Find(&user)
	db.First(&user, `name =?`, "Tim")
	fmt.Println(user)
	//改
	db.Where("age<?", "20").Update("age", 25)
	//删
	db.Where("age<?", "20").Delete(&User{})
	defer db.Close()
}
