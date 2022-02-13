//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

const (
	newborn    = 0
	toddler    = 1
	child 	   = 4
	teenager   = 13
	adult      = 18
)

func main() {
	age := child;
	switch age {
	case adult:
		fmt.Println("adult")
	case teenager:
		fmt.Println("teenager")
	case child:
		fmt.Println("child")
	case toddler:
		fmt.Println("toddler")
	default:
		fmt.Println("newborn")
	}
	
}
