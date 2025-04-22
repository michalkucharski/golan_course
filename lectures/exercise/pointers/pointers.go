//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Activated    = true
	Desactivated = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Activated
}

func inactivate(tag *SecurityTag) {
	*tag = Desactivated
}

func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		fmt.Println("item:", items[i])
		inactivate(&items[i].tag)

	}

}

func main() {

	items := []Item{
		{name: "T-shirt", tag: true},
		{name: "Short", tag: true},
		{name: "Cap", tag: true},
		{name: "Bra", tag: true},
	}

	fmt.Println("State without any changes", items)

	inactivate(&items[2].tag)

	fmt.Println(items)

	checkout(items)

	fmt.Println("All items desactivated", items)

}
