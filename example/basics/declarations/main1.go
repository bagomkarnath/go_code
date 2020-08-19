package main

import "fmt"

// Title constant
const Title string = "Person Details"

// Country package variable
var Country string = "USA"

func main() {
	var fname, lname string = "omkar", "bag"
	fmt.Println("First name : ", fname)
	fmt.Println("Last name : ", lname)
	var age int = 32
	fmt.Println("Age : ", age)
}
