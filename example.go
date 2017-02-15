package main

import "fmt"

func updateValue(someVal *int) {
	*someVal = *someVal + 100
}

func main() {
	val := 1000
	updateValue(&val)
	fmt.Println("val:", val)
}
