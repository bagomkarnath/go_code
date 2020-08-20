package main

import (
	"fmt"

	"github.com/bagomkarnath/go_code/example/nestedpackage/domain/student"
	"github.com/bagomkarnath/go_code/example/nestedpackage/domain/student/btech"
)

func main() {
	fmt.Println("student.Name : ", student.Name)
	fmt.Println("student.btech.StreamName : ", btech.StreamName)
}
