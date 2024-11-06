package grades

// 声明变量 students
var students []Student

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "John", // 修正为 FirstName
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Math",
					Type:  GradeQuiz,
					Score: 80,
				},
				{
					Title: "Science",
					Type:  GradeExam,
					Score: 90,
				},
				{
					Title: "English",
					Type:  GradeAssignment,
					Score: 75,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jane", // 修正为 FirstName
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Math",
					Type:  GradeQuiz,
					Score: 70,
				},
				{
					Title: "Science",
					Type:  GradeExam,
					Score: 80,
				},
				{
					Title: "English",
					Type:  GradeAssignment,
					Score: 85,
				},
			},
		},
	}
}
