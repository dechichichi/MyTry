package grades

import "fmt"

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

func (s Student) Average() float32 {
	var result float32
	for _, g := range s.Grades {
		result += g.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var students []Students

func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("student with id %d not found", id)
}
