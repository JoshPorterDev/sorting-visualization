package algorithms

import (
	"github.com/joshporter3/sorting-visualization/draw"
)

func BubbleSort(arr []int) {
	var swapped bool
	for i := 0; i < len(arr); i++ {
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				draw.Draw(arr)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
