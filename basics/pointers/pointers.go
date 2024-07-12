package main

import "fmt"

func main() {
	age := 32

	// * signifies pointer
	var agePointer *int
	// & is the pointer
	agePointer = &age
	// * gets value from pointer
	fmt.Println("age: ", *agePointer)

	// pointer overrides initial age
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
