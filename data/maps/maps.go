package main

import "fmt"

//maps are dynamic no need for slicing
func main() {
	// key and value types
	websites := map[string]string{
		"Google": "google.com",
		"AWS":    "aws.com",
	}
	fmt.Println(websites)
	// only print value of key
	fmt.Println(websites["AWS"])
	// add url
	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)
	// delete key
	delete(websites, "Google")
	fmt.Println(websites)
}
