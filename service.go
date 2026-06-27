package main

import "fmt"

var students []Student

// Add Student
func AddStudent(student Student) {
	students = append(students, student)
	fmt.Println("Student Added Successfully")
}

// Search Student
func SearchStudent(id int) {

	for _, student := range students {

		if student.ID == id {

			fmt.Println("\nStudent Found")
			fmt.Println("------------------------")
			fmt.Println("ID     :", student.ID)
			fmt.Println("Name   :", student.Name)
			fmt.Println("Course :", student.Course)

			return
		}
	}

	fmt.Println("Student Not Found")
}
