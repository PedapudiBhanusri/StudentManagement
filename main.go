package main

func main() {

	AddStudent(Student{
		ID:     1,
		Name:   "Bhanu",
		Course: "Go",
	})

	AddStudent(Student{
		ID:     2,
		Name:   "Rahul",
		Course: "Java",
	})

	AddStudent(Student{
		ID:     3,
		Name:   "Priya",
		Course: "Python",
	})

	SearchStudent(2)
}
