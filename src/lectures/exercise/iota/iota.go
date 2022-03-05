//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

func (o Operation) calculate(v1 int, v2 int) int {
	switch o {
	case Add:
		return v1+v2
	case Sub:
		return v1-v2
	case Mul:
		return v1*v2
	case Div:
		return v1/v2
	}
	panic("unhandled operation")
}

func main() {
	add := Add
	sub := Sub
	mul := Mul
	div := Div
	
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
