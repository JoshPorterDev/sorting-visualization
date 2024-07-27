package array

import (
	"math/rand"
)

func ShuffleArray(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func GenerateArray(n int) []int {
	arr := make([]int, 200)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	ShuffleArray(arr)
	return arr
}
