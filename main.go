package main

import "fmt"

func main() {
	s := []int{1, 4, 7, 11, 67, 90}
	item := 90
	res := binarySearch(s, item)
	fmt.Println("Iterations:", res.counter, "Index", res.index)

}

type result struct {
	index   int
	counter int
}

func binarySearch(s []int, item int) result {
	low := 0
	high := len(s) - 1
	counter := 1

	for low <= high {
		mid := (low + high) / 2
		if s[mid] == item {
			return result{index: mid, counter: counter}
		}
		if s[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
		counter++
	}

	return result{index: -1, counter: counter}
}
