package main

import(
	"fmt"
)

func main()  {
	//constant = var yang tidak bisa berubah 

	const fName = "Mario"
	const lName = "Aprilnino"
	fmt.Println(fName)
	fmt.Println(lName)
	//if we use const we cant change the variable

	//decalre multiple constant
	const(
		firstName	= "Mario"
		lastName	= "Aprlnino"
		age		= 23
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}