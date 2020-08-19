package main

import (
	"fmt"
)

var myMap map[string]string

func init() {
	myMap = make(map[string]string)
}

func show(smap map[string]string) {
	fmt.Println("Length of Map is : ", len(smap))
	for k, v := range smap {
		fmt.Printf("key : %s > value : %s \n", k, v)
	}
}

func add(k, v string) {
	if _, ok := myMap[k]; ok {
		fmt.Println("Key already exists")
	} else {
		myMap[k] = v
	}

}

func getAll() map[string]string {
	return myMap
}

func get(k string) string {
	var value string
	if _, ok := myMap[k]; ok {
		value = myMap[k]
	} else {
		fmt.Println("Key doesn't exists")
	}
	return value
}

func update(k, v string) {
	if _, ok := myMap[k]; ok {
		fmt.Printf("Updating k > %s : v > %s \n", k, v)
		myMap[k] = v
	} else {
		fmt.Println("Key doesn't exists")
	}
}

func del(k string) {
	if _, ok := myMap[k]; ok {
		fmt.Printf("Deleting k > %s : v > %s \n", k, myMap[k])
		delete(myMap, k)
	} else {
		fmt.Println("Key doesn't exists")
	}
}

func main() {
	add("1", "Apple")
	add("2", "Orange")
	add("3", "Banana")
	add("4", "Grapes")
	add("5", "Mango")
	show(myMap)
	fmt.Println("Value at 4 : ", get("4"))
	update("5", "Sweet Mango")
	show(myMap)
	del("1")
	show(myMap)
	add("1", "Pine apple")

	testMap := getAll()

	show(testMap)

	update("6", "Test")

	del("7")
}
