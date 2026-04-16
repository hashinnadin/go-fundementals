package main

import (
	"errors"
	"fmt"
)

// Struct
type Student struct {
	Name string
	Age  int
}

// Slice to store students
var students []Student

// Add Student
func AddStudent(name string, age int) {
	students = append(students, Student{Name: name, Age: age})
	fmt.Println("Student added successfully")
}

// List Students
func ListStudents() {
	if len(students) == 0 {
		fmt.Println("No students found")
		return
	}
	for i, s := range students {
		fmt.Printf("%d. Name: %s, Age: %d\n", i, s.Name, s.Age)
	}
}

// Update Student (using pointer)
func UpdateStudent(index int, name string, age int) error {
	if index < 0 || index >= len(students) {
		return errors.New("invalid student index")
	}

	// pointer usage
	student := &students[index]
	student.Name = name
	student.Age = age

	return nil
}

// Delete Student
func DeleteStudent(index int) error {
	if index < 0 || index >= len(students) {
		return errors.New("invalid student index")
	}

	students = append(students[:index], students[index+1:]...)
	return nil
}

// Menu
func main() {
	for {
		fmt.Println("\n--- Student Management ---")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var name string
			var age int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter age: ")
			fmt.Scan(&age)

			AddStudent(name, age)

		case 2:
			ListStudents()

		case 3:
			var index, age int
			var name string

			fmt.Print("Enter index: ")
			fmt.Scan(&index)
			fmt.Print("Enter new name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new age: ")
			fmt.Scan(&age)

			err := UpdateStudent(index, name, age)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Student updated")
			}

		case 4:
			var index int
			fmt.Print("Enter index: ")
			fmt.Scan(&index)

			err := DeleteStudent(index)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Student deleted")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
