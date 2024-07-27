package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joshporter3/sorting-visualization/algorithms"
	"github.com/joshporter3/sorting-visualization/array"
)

func HandleInput(arr []int) {
	if rl.IsKeyReleased(rl.KeyQ) {
		algorithms.QuickSort(arr, 0, 199)
	}
	if rl.IsKeyReleased(rl.KeyB) {
		algorithms.BubbleSort(arr)
	}
	if rl.IsKeyReleased(rl.KeyM) {
		algorithms.MergeSort(arr, 0, 199)
	}

	if rl.IsKeyReleased(rl.KeyS) {
		array.ShuffleArray(arr)
	}
}
