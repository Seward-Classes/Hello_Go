package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	values := randomIntList(10)
	fmt.Println("Random list:", values)
	quickSort(values)
	fmt.Println("Sorted list:", values)
}

func randomIntList(length int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := make([]int, length)
	for i := range values {
		values[i] = rng.Intn(100)
	}
	return values
}

func quickSort(values []int) {
	quickSortRange(values, 0, len(values)-1)
}

func quickSortRange(values []int, low, high int) {
	if low >= high {
		return
	}

	pivot := values[low+(high-low)/2]
	less, current, greater := low, low, high
	for current <= greater {
		switch {
		case values[current] < pivot:
			values[less], values[current] = values[current], values[less]
			less++
			current++
		case values[current] > pivot:
			values[current], values[greater] = values[greater], values[current]
			greater--
		default:
			current++
		}
	}

	quickSortRange(values, low, less-1)
	quickSortRange(values, greater+1, high)
}
