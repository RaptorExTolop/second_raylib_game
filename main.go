package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running            bool  = true
	bkgcolour                = rl.SkyBlue
	screenWidth        int32 = 1280
	screenHeight       int32 = 720
	collidingWithFloor bool  = false

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
	playerJumpHeight            int
	playerLastKnownY            int32

	// gravity
	gravity      float32
	falling_time float32

	// player animation
	playerRunningFrame = 0
	playerJumpFrame    = 0
	playerLastDir      = 0

	frameCountPerSec = 0
)

const ()

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerJumping = true
		//fmt.Println("jumping")
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir -= 1
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir += 1
	}
	playerDir = int(rl.Clamp(float32(playerDir), -1, 1))
	/*if rl.IsKeyPressed(rl.KeyL) {
		collidingWithFloor = !collidingWithFloor
	}*/
	if playerDir == -1 || playerDir == 1 {
		playerLastDir = playerDir
	}
}

func update() {
	running = !rl.WindowShouldClose()
	frameCountPerSec++
	if (frameCountPerSec % 6) == 0 {
		playerRunningFrame++
	}
	if frameCountPerSec >= 60 {
		frameCountPerSec = 0
	}
	if playerRunningFrame > 7 {
		playerRunningFrame = 0
	}
	if (frameCountPerSec % 3) == 0 {
		playerJumpFrame++
	}
	if playerJumpFrame > 7 {
		playerJumpFrame = 0
	}
	//fmt.Println(frameCountPerSec)
	// gravity & jumping
	if !collidingWithFloor {
		falling_time += 0.0166
		falling_time = rl.Clamp(falling_time, -1, 15)
		playerDest.Y += falling_time * gravity
	} else if collidingWithFloor && playerJumping {
		falling_time = 0
		playerDest.Y -= float32(playerJumpHeight)
	}

	// player controls

	// player idle animaiton
	if frameCountPerSec%5 == 0 {
		if playerLastDir == -1 {
			playerSrc.Y = 8 * 56
		} else if playerLastDir == 1 {
			playerSrc.Y = 0
		} else if playerLastDir == 0 {
			playerSrc.Y = 0
		}
		playerSrc.X = float32(56 * (frameCountPerSec / 10))
	}
	// player moving & running animation
	if playerMoving {
		if playerDir == -1 {
			//fmt.Println("left")
			playerDest.X -= playerSpeed
			playerSrc.Y = 7 * 56
			playerSrc.X = float32(56 * playerRunningFrame)
		} else if playerDir == 1 {
			//fmt.Println("right")
			playerDest.X += playerSpeed
			playerSrc.X = float32(56 * playerRunningFrame)
			playerSrc.Y = 2 * 56
		}
	}
	if playerDest.Y >= float32(screenHeight)-168 {
		playerDest.Y = float32(screenHeight) - 168
		collidingWithFloor = true
	} else {
		collidingWithFloor = false
	}

	if float32(playerLastKnownY) > playerDest.Y {
		fmt.Println("up")
		if playerLastDir == -1 {
			playerSrc.Y = 9 * 56
			playerSrc.X = float32(56 * playerJumpFrame)
		} else if playerLastDir == 1 {
			playerSrc.Y = 3 * 56
			playerSrc.X = float32(56 * playerJumpFrame)
		}
	} else if float32(playerLastKnownY) < playerDest.Y {
		fmt.Println("down")
		if playerLastDir == -1 {
			playerSrc.Y = 10 * 56
			playerSrc.X = float32(56 * playerJumpFrame)
		} else if playerLastDir == 1 {
			playerSrc.Y = 4 * 56
			playerSrc.X = float32(56 * playerJumpFrame)
		}
	}

	playerLastKnownY = int32(playerDest.Y)

	playerDir = 0
	playerMoving = false
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
	rl.UnloadTexture(bkgl1)
	rl.UnloadTexture(bkgl2)
	rl.UnloadTexture(bkgl3)
	rl.UnloadTexture(playerSprite)

	rl.CloseWindow()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	gravity = 35
	playerJumpHeight = 56 * 4

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
	//float32(screenHeight)-168
	playerSpeed = 4.0
}

func main() {

	for running {
		input()
		update()
		draw()
	}
	quit()
}
