package main

import "fmt"

func main() {
	var studentName, grade, isPassed = "Jhon Doe", 77, true

	studentName2, grade2, isPassed2 := "Jhon Doe2", 77, true

	var (
		studentName3 string = "Jhon Doe3"
		grade3       int    = 77
		isPassed3    bool   = true
	)

	surname := "ST" //Declaration
	surname = "STT" //Assign
	
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
	fmt.Println("------------------")
	fmt.Println(studentName2)
	fmt.Println(grade2)
	fmt.Println(isPassed2)
	fmt.Println("------------------")
	fmt.Println(studentName3)
	fmt.Println(grade3)
	fmt.Println(isPassed3)
	fmt.Println("------------------")
	fmt.Println(surname)

}