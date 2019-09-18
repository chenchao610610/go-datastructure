package main

/*
 * Quick sort - https://en.wikipedia.org/wiki/Quicksort
 */

import "fmt"
import "math/rand"

import utils "github.com/0xAX/go-algorithms"

func quick_sort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	//fmt.Println(len(arr))
	median := arr[rand.Intn(len(arr))]
	counter := 0
	for _, item := range arr {
		if item == median {
			counter = counter + 1
		}
	}

	if len(arr) == counter {
		return arr
	}

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	//middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			low_part = append(low_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	fmt.Println(low_part, "==", median, "==", high_part)

	low_part = quick_sort(low_part)
	high_part = quick_sort(high_part)

	//low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", quick_sort(arr))
}
