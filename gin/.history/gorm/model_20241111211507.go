package dbdata

import "github.com/jinzhu/gorm"

type Man struct {
	gorm.Model
	//加了一个primary_key主键
	//自定义列名称user_name
	//自定义列类型varchar(255)
	Name string `gorm:"primary_key;column:'user_name';type:'varchar(255)'"`
}

func (m Man) TableName() string {
	if m.Name == "" {
		return "qu_man"
	} else {
		return "qu_man_" + m.Name
	}
}

func Ia() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	db.AutoMigrate(&Man{})
	if err != nil {
		panic(err)
	}
}
