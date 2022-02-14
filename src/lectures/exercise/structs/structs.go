//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rect struct {
	Height int
	Width int
}

func rectPerimeter(rect Rect) int {
	return 2 * (rect.Height + rect.Width)
}

func rectArea(rect Rect) int {
	return rect.Height * rect.Width
}

func main() {

	fmt.Println(rectPerimeter(Rect{2,3}))
	fmt.Println(rectArea(Rect{2,4}))

}
