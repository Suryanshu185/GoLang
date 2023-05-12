package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":       4.99,
		"pie":        7.99,
		"salad":      6.99,
		"cheesecake": 8.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key types
	phone := map[int]string{
		7006347045: "suryanshu",
		7006347044: "sehajal",
		9419909074: "mama",
		7889390439: "panchi",
	}
	fmt.Println(phone)
	fmt.Println(phone[7006347044])

	phone[9419909074] = "suryanshu"
	phone[7889390439] = "sehajal"

	fmt.Println(phone)
}
