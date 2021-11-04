package main

import (
	"fmt"
)

func BinarySearch(list []int, item int) (pos int) {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + high
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}

func main() {
	myList := []int{0, 1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(myList, 3))
}
