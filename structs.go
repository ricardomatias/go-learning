package main

import "fmt"

type person struct {
	name string
	age  int
}

func p(people ...person) {
	fmt.Println(people)
}

func main() {
	p(person{"Bob", 20})

	p(person{age: 40, name: "foo"})

	p(person{age: 40})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "hello", age: 10000}

	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)

	sp.age = 21
	fmt.Println(sp)
}
