package main

import "fmt"

func main() {
	size := 10
	var numbers []int

	for i := 0; i <= size; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, " is EVEN NUMBER ")
		} else {
			fmt.Println(num, " is ODD NUMBER ")
		}
	}
}
