package main

import "fmt"

func main() {
	//fmt.Println("Hello World!")
	var input int
	fmt.Println("kuch to likh bc:")
	fmt.Scanf("%v", &input)

	s := make([]int, input)

	for i := 0; i < input; i++ {
		fmt.Scanf("%v", &s[i])
	}

	/*for _, value := range s {
	//	fmt.Println(value)
	//	sum += value
	*/

	for j := input - 1; j >= 0; j-- {
		fmt.Println(s[j])
	}
}
