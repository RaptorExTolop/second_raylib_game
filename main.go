package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running      bool  = true
	bkgcolour          = rl.SkyBlue
	screenWidth  int32 = 1280
	screenHeight int32 = 720
	gravity      float32

	// background stuff
	bkgl1     rl.Texture2D
	bkgl1Src  rl.Rectangle
	bkgL1Dest rl.Rectangle

	bkgl2     rl.Texture2D
	bkgl2Src  rl.Rectangle
	bkgL2Dest rl.Rectangle

	bkgl3     rl.Texture2D
	bkgl3Src  rl.Rectangle
	bkgL3Dest rl.Rectangle

	// player
	playerMoving, playerJumping bool
	playerDir                   int
	playerSprite                rl.Texture2D
	playerSrc                   rl.Rectangle
	playerDest                  rl.Rectangle
	playerSpeed                 float32
)

const ()

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerJumping = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir += 1
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir -= 1
	}
	playerDir = int(rl.Clamp(float32(playerDir), -1, 1))
}

func update() {
	running = !rl.WindowShouldClose()
	if !playerJumping {
		playerDest.Y += gravity
	} else {
		playerDest.Y -= 5
	}

	if playerMoving {
		if playerDir == 1 {
			fmt.Println("left")
			playerDest.X -= playerSpeed
		} else if playerDir == -1 {
			fmt.Println("right")
			playerDest.X += playerSpeed
		}
	}
	playerDir = 0
	playerMoving = false
	if playerJumping {
		fmt.Println("jumping")
	}
	playerJumping = false
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	// draw the three background layers
	rl.DrawTexturePro(bkgl1, bkgl1Src, bkgL1Dest, rl.NewVector2(0, 0), 0, rl.RayWhite)
	rl.DrawTexturePro(bkgl2, bkgl2Src, bkgL2Dest, rl.NewVector2(0, 0), 0, rl.RayWhite)
	rl.DrawTexturePro(bkgl3, bkgl3Src, bkgL3Dest, rl.NewVector2(0, 0), 0, rl.RayWhite)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.RayWhite)

	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	gravity = 3

	// background inits
	bkgl1 = rl.LoadTexture("res/background/background_layer_1.png")
	bkgL1Dest = rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))
	bkgl1Src = rl.NewRectangle(0, 0, 320, 180)

	bkgl2 = rl.LoadTexture("res/background/background_layer_2.png")
	bkgL2Dest = rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))
	bkgl2Src = rl.NewRectangle(0, 0, 320, 180)

	bkgl3 = rl.LoadTexture("res/background/background_layer_3.png")
	bkgL3Dest = rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))
	bkgl3Src = rl.NewRectangle(0, 0, 320, 180)

	// player inits
	playerSprite = rl.LoadTexture("res/character/char_blue.png")
	playerSrc = rl.NewRectangle(0, 0, 56, 56)
	playerDest = rl.NewRectangle(0, 0, 168, 168)
	playerSpeed = 3.0
}

func main() {

	for running {
		input()
		update()
		draw()
	}
	quit()
}
