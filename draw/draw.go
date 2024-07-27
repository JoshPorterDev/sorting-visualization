package draw

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawArray(arr []int) {
	for index, element := range arr {
		rect_offset := index * 4
		rect_height := element * 4
		rl.DrawRectangle(int32(rect_offset), int32(rl.GetScreenHeight()-rect_height), 4, int32(rect_height), rl.Blue)
	}
}

func Draw(arr []int) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	rl.DrawText("Sorting Visualization", 0, 0, 20, rl.Blue)

	drawArray(arr)

	rl.EndDrawing()
}
