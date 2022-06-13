package main

import "fmt"

var packVar = "Package Scope"

func main() {
	funcVar := "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)
	myFunc()
}
func myFunc()  {
	fmt.Println(packVar)
}