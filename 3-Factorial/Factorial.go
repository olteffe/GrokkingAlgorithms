package main

import "fmt"

func factorialRecursive(number uint64) uint64 {
	if number >= 1 {
		return number * factorialRecursive(number-1)
	}
	return 1
}

func factorialIterative(number uint64) uint64 {
	var result uint64 = 1
	if number > 1 {
		for i := 1; i <= int(number); i++ {
			result *= uint64(i)
		}
	}
	return result
}

func main() {
	fmt.Println(factorialRecursive(5))
	fmt.Println(factorialIterative(5))
}
