package main

import "fmt"

func main() {
	numChecker()
}

func numChecker() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers)

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
