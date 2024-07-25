package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 820, "Sorting Visualization")

	rl.SetTargetFPS(60)

	arr := generateArray(200)

	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Sorting Visualization", 0, 0, 20, rl.Blue)
		drawArray(arr)

		rl.EndDrawing()
	}
}

func shuffleArray(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func generateArray(n int) []int {
	arr := make([]int, 200)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	shuffleArray(arr)
	return arr
}

func drawArray(arr []int) {
	for index, element := range arr {
		rect_offset := index * 4
		rect_height := element * 4
		rl.DrawRectangle(int32(rect_offset), int32(rl.GetScreenHeight()-rect_height), 4, int32(rect_height), rl.Blue)
	}
}
