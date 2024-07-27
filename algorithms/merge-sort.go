package algorithms

import "github.com/joshporter3/sorting-visualization/draw"

func MergeSort(arr []int, left int, right int) {
	if left < right {
		mid := (right + left) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left int, mid int, right int) {
	left_length := mid - left + 1
	right_length := right - mid
	left_arr := make([]int, left_length)
	right_arr := make([]int, right_length)

	for i := 0; i < left_length; i++ {
		left_arr[i] = arr[left+i]
	}
	for i := 0; i < right_length; i++ {
		right_arr[i] = arr[mid+i+1]
	}

	i := 0
	j := 0
	k := left
	for i < left_length && j < right_length {
		if left_arr[i] < right_arr[j] {
			arr[k] = left_arr[i]
			draw.Draw(arr)
			i++
		} else {
			arr[k] = right_arr[j]
			draw.Draw(arr)
			j++
		}
		k++
	}
	for i < left_length {
		arr[k] = left_arr[i]
		draw.Draw(arr)
		k++
		i++
	}
	for j < right_length {
		arr[k] = right_arr[j]
		draw.Draw(arr)
		k++
		j++
	}
}
