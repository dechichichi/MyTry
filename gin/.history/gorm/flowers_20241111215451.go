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
	id := IDCard{
		Number: 123456,
	}
	stu := Student{
		StudentName: "Tom",
		IDCard:      id,
	}
	cla := Class{
		Students: []Student{stu},
	}
	tea := Teacher{
		TeacherName: "Jane",
		Students:    []Student{stu},
	}
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
	_ = db.Create(&id).Error
	_ = db.Create(&stu).Error
	_ = db.Create(&cla).Error
	_ = db.Create(&tea).Error
}
