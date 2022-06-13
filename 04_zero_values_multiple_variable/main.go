package main

import "fmt"

func main() {
	var (
		name string = "Ahmet"
		secondName string = "Emin"
		age = 30
		isMale bool = true
		weight float32 = 69.5
		height = 178.7
	)
	surname := "ST" //Short declaration

	var name2, age2, isMale2, weight2, height2= "Chloe", 25, false, 49.4, 167.8
	//name2, age2, isMale2, weight2, height2 := "Chloe", 25, false, 49.4, 167.8 

	var defaultString string //---> ""
	var defaultInt int //---> 0

	fmt.Println(name);
	fmt.Println(secondName);
	fmt.Println(surname);
	fmt.Println(age);
	fmt.Println(isMale);
	fmt.Println(weight);
	fmt.Println(height);
	fmt.Println(name2);
	fmt.Println(age2);
	fmt.Println(isMale2);
	fmt.Println(weight2);
	fmt.Println(height2);
	fmt.Println(defaultString);
	fmt.Println(defaultInt);

}