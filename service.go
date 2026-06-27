package main

import "fmt"

var students []Student

func AddStudent(student Student) {

	students = append(students, student)

	fmt.Println("Student Added Successfully")
}
