package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/joshporter3/sorting-visualization/array"
	"github.com/joshporter3/sorting-visualization/draw"
)

func main() {
	rl.InitWindow(800, 820, "Sorting Visualization")

	rl.SetTargetFPS(60)

	arr := array.GenerateArray(200)

	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {

		draw.Draw(arr)
	}
}
