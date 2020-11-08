package main

import "fmt"

const (
	Ascending = iota
	Descending
)

func main() {
	arr := []int{3,4,1,2,6,5,10,9,8,11,23,45,22,89,43,49,34}
	arr = SortSliceInt(arr, 1)
	fmt.Println(arr)
}

func SortSliceInt(arr []int, t int8) []int {
	result := make([]int, len(arr))
	index := len(arr)-1
	bubble(arr, result, index, t)
	return result
}

func bubble(arr []int, result []int, index int, t int8){
	greatestVal := arr[0]
	ix := 0
	for j, val := range arr {
		switch t {
		case Ascending:
			if val > greatestVal {
				greatestVal = val
				ix = j
			}
		case Descending:
			if val < greatestVal {
				greatestVal = val
				ix = j
			}
		}
	}
	arr = removeVal(ix, arr)
	result[index] = greatestVal
	index -= 1
	if len(arr) > 1 {
		bubble(arr, result, index, t)
	} else {
		result[0] = arr[0]
	}
}

func removeVal(index int, arr []int) []int {
	var result []int
	result = append(arr[:index], arr[index+1:]...)
	return result
}