package algorithms

import "github.com/joshporter3/sorting-visualization/draw"

func QuickSort(arr []int, left int, right int) {
	if left < right {
		partition := partition(arr, left, right)
		QuickSort(arr, left, partition-1)
		QuickSort(arr, partition+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	x := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= x {
			i += 1
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
			draw.Draw(arr)
		}
	}
	temp := arr[right]
	arr[right] = arr[i+1]
	arr[i+1] = temp
	draw.Draw(arr)
	return i + 1
}
