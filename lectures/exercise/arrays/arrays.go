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

type Shopitem struct {
	name  string
	price int
}

func checkShoplist(shopList [4]Shopitem) {
	totalPrice := 0

	for i := 0; i < len(shopList); i++ {
		shopList := shopList[i]

		totalPrice = shopList.price + totalPrice
	}

	fmt.Println("Last item of list: ", shopList[len(shopList)-1])
	fmt.Println("Total number of items: ", len(shopList))
	fmt.Println("Total cost of items:", totalPrice)
}

func main() {

	shopList := [4]Shopitem{
		{name: "Banana", price: 4},
		{name: "Una patata", price: 1},
		{name: "Fish", price: 10},
	}

	checkShoplist(shopList)

	fmt.Println("Adding the last product")

	shopList[3] = Shopitem{"bread", 5}

	checkShoplist(shopList)

}
