package dbdata

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	Students []Student
}

type Student struct {
	gorm.Model
	ClassID     uint
	StudentName string
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
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers"`
	StudentID   uint
}

func Go() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stu := Student{
		StudentName: "Tom",
	}
	cla := Class{}
	id := IDCard{
		Number: 123456,
	}
	tea := Teacher{
		StudentID: 1,
	}
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
}
