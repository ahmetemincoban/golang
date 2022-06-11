//Package clause
package main

//Import statement
import "fmt" //---> Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

//My Code
func main()  {
	
	var name string = "Ahmet" //variable | name of variable | static type | variable value

	var age int
	age = 30 //Set value in integer age variable

	var isMale bool
	isMale=true

	secondName:="Emin"

	fmt.Println("Hello", name, secondName)
	fmt.Println(age)
	fmt.Println(isMale)
}

/*
Output:
	Hello Ahmet Emin
	30
	true
*/