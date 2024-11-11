package dbdata

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	Students []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers"`
	TeacherID   uint
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Number    int
}

type Teacher struct {
	gorm.Model
	Students  []Student `gorm:"many2many:student_teachers"`
	StudentID uint
}

func Go() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stu := Student{}
	cla := Class{}
	id := IDCard{}
	tea := Teacher{}
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
}
