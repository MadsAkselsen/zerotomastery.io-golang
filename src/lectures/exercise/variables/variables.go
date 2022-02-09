//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor = "blue"
	fmt.Println("favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	year, ageInYears := 1991, 31
	fmt.Println("year is", year, "and age us", ageInYears)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstName = "Mads"
		lastName = "Akselsen"
	)
	fmt.Println(firstName, lastName)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var age int
	age = 11315/365
	fmt.Println("age is", age)
}
