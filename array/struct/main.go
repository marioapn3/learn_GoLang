package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 person
	p1.name = "Mario"
	p1.age = 30
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	//declare struct with key value
	p2 := person{name: "Aprilnino", age: 30}
	fmt.Println(p2.name)
	fmt.Println(p2.age)

	//declare struct without key value
	p3 := person{"Mario Aprilnino", 30}
	fmt.Println(p3.name)
	fmt.Println(p3.age)

	//declare struct with pointer
	p4 := &person{"Mario Aprilnino", 30}
	fmt.Println(p4.name)
	fmt.Println(p4.age)
}
