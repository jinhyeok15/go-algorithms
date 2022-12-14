package main

import (
	"fmt"
)

func main() {
	givenArr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Given Slice")
	fmt.Println(givenArr)

	sort(givenArr, 0, len(givenArr) - 1)
	fmt.Println("\nSorted slice")
	fmt.Println(givenArr)
}

func merge(arr []int, left int, mid int, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid
	
	L, R := make([]int, 0), make([]int, 0)

	for i:=0; i<leftSize; i++ {
		L = append(L, arr[left + i])
	}

	for j:=0; j<rightSize; j++ {
		R = append(R, arr[mid + 1 + j])
	}

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if (L[i] <= R[j]) {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < leftSize {
		arr[k] = L[i]
		i++
		k++
	}

	for j < rightSize {
		arr[k] = R[j]
		j++
		k++
	}
}

func sort(arr []int, l int, r int) {
	if l < r {
		m := l + (r - l) / 2

		sort(arr, l, m)
		sort(arr, m + 1, r)

		merge(arr, l, m, r)
	}
}
