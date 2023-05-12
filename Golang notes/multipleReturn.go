package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	first1, second1 := getInitials("sehajal sharma")
	fmt.Println(first1, second1)

	first2, second2 := getInitials("suryanshu gupta")
	fmt.Println(first2, second2)

	first3, second3 := getInitials("sehajal")
	fmt.Println(first3, second3)

}
