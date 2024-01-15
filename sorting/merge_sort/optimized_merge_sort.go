package main

import "fmt"

func OptimizedMergeSort(array []int) ([]int, int) {
	if len(array) <= 1 {
		return array, 0
	}

	auxArr := make([]int, len(array))
	count := 0
	var inversionCount *int = &count

	copy(auxArr, array)
	mergeSortWithAux(array, 0, len(array)-1, auxArr, inversionCount)

	return array, *inversionCount
}

func mergeSortWithAux(mainArr []int, start, end int, auxArr []int, inversionCount *int) {
	if start == end {
		return
	}

	mid := (end-start)/2 + start

	mergeSortWithAux(auxArr, start, mid, mainArr, inversionCount)
	mergeSortWithAux(auxArr, mid+1, end, mainArr, inversionCount)

	doMerge(mainArr, start, mid, end, auxArr, inversionCount)
}

func doMerge(mainArr []int, start, mid, end int, auxArr []int, inversionCount *int) {
	k, i, j := start, start, mid+1
	for i <= mid || j <= end {
		if i > mid {
			mainArr[k] = auxArr[j]
			j++
			k++
			continue
		}

		if j > end {
			mainArr[k] = auxArr[i]
			i++
			k++
			continue
		}

		if auxArr[j] < auxArr[i] {
			mainArr[k] = auxArr[j]
			j++
			*inversionCount += mid - i + 1
		} else {
			mainArr[k] = auxArr[i]
			i++
		}

		k++
	}
}

func optimizedMergeSortMain() {
	arr := []int{2, 3, 3, 1, 9, 5, 6}

	fmt.Println(OptimizedMergeSort(arr))
}
