package main

import (
	"fmt"
)

// Add func
func Add(x, y int) int {
	return x + y
}

// Substract func
func Substract(x, y int) int {
	return x - y
}

func main() {
	x, y := 30, 10
	resultAdd := Add(x, y)
	resultSub := Substract(x, y)
	fmt.Println("Add : ", resultAdd)
	fmt.Println("Sub : ", resultSub)
}
