package main

import "fmt"

func updateName(x *string) {
	*x = "sehajal"
}

func main() {

	name := "suryanshu"

	// fmt.Println("the address of the name is", &name)

	m := &name
	fmt.Println("the name was", name)
	updateName(m)
	// fmt.Println("memory address", m)
	// fmt.Println("value at memory address", *m)
	fmt.Println("the name is", name)

}
