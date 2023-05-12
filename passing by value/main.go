package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "sehajal"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types :- strings , ints , bools , floats , arrays , structs
	name := "suryanshu"

	// updateName(name) //this just changes a copy of the original name
	// to change the original value do this

	name = updateName(name)

	fmt.Println(name)

	// group B types :- slices , maps functions
	menu := map[string]float64{
		"pie":       3.99,
		"ice-cream": 5.95,
	}
	updateMenu(menu)
	fmt.Println(menu)

}
