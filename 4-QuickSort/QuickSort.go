package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type split struct {
	less    []int
	pivot   int
	greater []int
}

func (s *split) add(slice []int) {
	for _, value := range slice {
		if value < s.pivot {
			s.less = append(s.less, value)
		} else {
			s.greater = append(s.greater, value)
		}
	}
}

func quickSortStruct(array []int) []int {
	if len(array) < 2 {
		return array
	}
	rand.Seed(time.Now().Unix())
	pivotIndex := rand.Intn(len(array))
	s := split{less: make([]int, 0), pivot: array[pivotIndex], greater: make([]int, 0)}
	s.add(array[:pivotIndex])
	s.add(array[pivotIndex+1:])
	return append(append(quickSortStruct(s.less), s.pivot), quickSortStruct(s.greater)...)
}

func quickSortSlice(array []int) []int {
	var less, greater []int
	if len(array) < 2 {
		return array
	}
	rand.Seed(time.Now().Unix())
	pivotIndex := rand.Intn(len(array))
	pivot := array[pivotIndex]
	array = append(array[:pivotIndex], array[pivotIndex+1:]...)
	for _, number := range array {
		if number < pivot {
			less = append(less, number)
		} else {
			greater = append(greater, number)
		}
	}
	return append(append(quickSortSlice(less), pivot), quickSortSlice(greater)...)
}

// Note: in golang v1.17 used modified quick sort
func goSort(array []int) []int {
	sort.Ints(array)
	return array
}

func main() {
	fmt.Println(quickSortStruct([]int{10, 5, 2, 3}))
	fmt.Println(quickSortSlice([]int{10, 5, 2, 3}))
	fmt.Println(goSort([]int{10, 5, 2, 3}))
}
