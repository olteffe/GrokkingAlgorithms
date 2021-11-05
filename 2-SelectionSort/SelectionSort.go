package main

import (
	"fmt"
	"sort"
)

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] < smallest {
			smallest = arr[idx]
			smallestIndex = idx
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) (newArr []int) {
	if len(arr) == 0 {
		return []int{}
	}
	for idx := len(arr) - 1; idx >= 0; idx-- {
		smallestIndex := findSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}

func goSort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	arr2 := []int{5, 3, 6, 2, 10, 11}
	fmt.Println(selectionSort(arr))
	fmt.Println(goSort(arr2))
}
