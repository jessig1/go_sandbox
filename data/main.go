package main

import "fmt"

func main() {
	// dynamic array with initial length with max capacity of 5
	userNames := make([]string, 2, 5)

	userNames[0] = "Paul"
	userNames = append(userNames, "Jake")
	userNames = append(userNames, "Britt")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["Terraform"] = 4.6
	courseRatings["Rust"] = 4.1

	fmt.Println(courseRatings)
}
