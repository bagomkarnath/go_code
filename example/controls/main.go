package main

import (
	"fmt"
)

func sum5() int {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func main() {
	sum5 := sum5()
	fmt.Println("Sum : ", sum5)
}
