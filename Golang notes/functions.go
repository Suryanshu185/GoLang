package main

import (
	"fmt"
	"math"
)

func greeting(n string) {
	fmt.Printf("MORNINGGGGGG %v \n", n)
}
func byebye(n string) {
	fmt.Printf("BYE BYE %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}
func areaCircle(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// greeting("sehajal")
	// byebye("sehajal")

	cycleNames([]string{"suryanshu", "aloo-bukhara", "sehajal"}, greeting)
	cycleNames([]string{"suryanshu", "aloo-bukhara", "sehajal"}, byebye)

	area1 := areaCircle(5.5)
	area2 := areaCircle(10.5)
	fmt.Println(area1, area2)

	fmt.Printf("area1 is %0.2f and area2 is %0.2f ", area1, area2)

}
