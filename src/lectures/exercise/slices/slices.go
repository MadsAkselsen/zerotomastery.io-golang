//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

func showLine(line []Part) {
	fmt.Println("---", "NEW ROUND", "---")
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

type Part string 

func main() {
	// make() function is used to preallocate a slice
	// useful when number of elements is known, but their values are stil unknown
	// will save computational power
	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "Part 1"
	assemblyLine[1] = "Part 2"
	assemblyLine[2] = "Part 3"
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Part 4", "Part 5")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	showLine(assemblyLine)
}
