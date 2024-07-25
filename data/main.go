package main

import "fmt"

// type alias makes it easier to define multiple types
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// dynamic array with initial length with max capacity of 5
	userNames := make([]string, 2, 5)

	userNames[0] = "Paul"
	userNames = append(userNames, "Jake")
	userNames = append(userNames, "Britt")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["Terraform"] = 4.6
	courseRatings["Rust"] = 4.1

	courseRatings.output()

	//fmt.Println(courseRatings)

	//iterates through array
	for index, value := range userNames {
		fmt.Print("index: ", index)
		fmt.Println("value: ", value)
	}

	//iterate a map
	for key, value := range courseRatings {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
}
