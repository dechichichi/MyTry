package grades

func init(){
	students=[]Student{
	{ID: 1,
	FirsName: "John",
	LastName: "Doe",
	Grades: []Grade{
		{Title: "Math",
		Type:GradeQuiz,
		Score: 80,
		},
		{
			Title: "Science",
			Type: GradeExam,
			Score: 90,
		},
		{
			Title: "English",
			Type: GradeAssignment,
			Score: 75,
		},
	},

		}
	
	}
}