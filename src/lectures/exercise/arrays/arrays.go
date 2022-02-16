//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price int
}

func main() {
	ShoppingList := [4]Product{
		{name: "Oranges", price: 10},
		{name: "apples", price: 20},
		{name: "Cheese", price: 43},
	}

	fmt.Println("The last item on the list", ShoppingList[len(ShoppingList)-1])
	fmt.Println("The total number of items", len(ShoppingList))

	totalPrice := 0
	for i := 0; i < len(ShoppingList); i++ {
		price := ShoppingList[i].price
		totalPrice += price
	}
	fmt.Println("The total cost of the items", totalPrice)

	ShoppingList[3].name = "Snacks"
	ShoppingList[3].price = 200

	fmt.Println("The last item on the list", ShoppingList[len(ShoppingList)-1])
}
