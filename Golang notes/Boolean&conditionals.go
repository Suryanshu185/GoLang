package main

import "fmt"

func main() {
	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 50)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is greater than 40")
	// }
	Names := []string{"suryanshu", "aloo-bukhara", "sehajal", "aloo", "bukhara"}

	for index, value := range Names {
		if index == 1 {
			fmt.Println("continuing at position \n", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at position \n", index)
			break
		}

		fmt.Printf("the value at position %v is %v \n", index, value)

	}

}
