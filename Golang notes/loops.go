package main

import "fmt"

func main() {

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is ", i)
	// }

	Names := []string{"suryanshu", "sehajal", "aloo-bukhara"}

	// for i := 0; i < len(Names); i++ {
	// 	fmt.Println(Names[i])
	// }
	// for index, value := range Names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	//
	for _, value := range Names {
		fmt.Printf("the value is %v \n", value)
	}

}
