package main

import "fmt"

// Container contains
type Container []interface{}

// Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

// Get gets an element from the container.
func (c *Container) Get(value int) interface{} {
	elem := (*c)[value]
	return elem
}

// The calling code does the type assertion when retrieving an element.
func main() {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(4)
	elem, ok := intContainer.Get(1).(int) // assert that the actual type is int
	if !ok {
		fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)
}
