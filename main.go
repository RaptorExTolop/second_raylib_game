package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	running   bool = true
	bkgcolour      = rl.SkyBlue
)

const ()

func input() {}

func update() {
	running = !rl.WindowShouldClose()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}

func init() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
}

func main() {

	for running {
		input()
		update()
		draw()
	}
	quit()
}
