package main

import "fmt"

func main() {

	var id int
	var name string
	var course string

	fmt.Println("===== Student Management =====")

	fmt.Print("Enter Student ID : ")
	fmt.Scan(&id)

	fmt.Print("Enter Student Name : ")
	fmt.Scan(&name)

	fmt.Print("Enter Course : ")
	fmt.Scan(&course)

	student := Student{
		ID:     id,
		Name:   name,
		Course: course,
	}

	AddStudent(student)
}
