package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 800, "Sorting Visualization")

	rl.SetTargetFPS(60)

	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Sorting Visualization", 0, 0, 20, rl.Blue)

		rl.EndDrawing()
	}
}
