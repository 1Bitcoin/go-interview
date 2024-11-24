package main

import "fmt"

type Person struct {
	name string
}

func changeName(person **Person) {
	*person = &Person{
		name: "Alice",
	}
}

func t4() {
	person := &Person{
		name: "Bob",
	}

	fmt.Println(person.name)
	changeName(&person)
	fmt.Println(person.name)
}
