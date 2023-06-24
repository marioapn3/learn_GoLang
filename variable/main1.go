package main
import (
	"fmt"
)
func main(){
	var name string 
	name = "This is Variable"
	fmt.Println(name)

	name = "I change this variable"
	fmt.Println(name)

	//cant change to different type data example change to int
	//, bool , and others

	//in golang we can make variable without data type example
	//if not declare the data type , Golang will determine the data type 
	var fName = "Mario"
	fmt.Println(fName)
	var age = 30
	fmt.Println(age)

	//we can make variable without use var , use :=
	lName := "Aprilnino"
	fmt.Println(lName)

	//declare multiple variable
	var(
		firstName = "Mario"
		lastName = "Aprilnino"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}