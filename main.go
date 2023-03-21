package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var myGame Game

	rl.InitWindow(ScreenWidth, ScreenHeight, "flappy")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	myGame.initGame()
	myGame.SpriteSheet = rl.LoadTexture("./assets/SpriteSheet.png")
	myGame.Player.FxJump = rl.LoadSound("./sounds/sfx_wing.wav")
	myGame.Player.FxHit = rl.LoadSound("./sounds/sfx_hit.wav")
	myGame.FxScore = rl.LoadSound("./sounds/sfx_point.wav")
	myGame.FxOver = rl.LoadSound("./sounds/sfx_die.wav")

	for !rl.WindowShouldClose() {

		update(&myGame)
		draw(&myGame)
	}
	rl.UnloadTexture(myGame.SpriteSheet)
	rl.UnloadSound(myGame.Player.FxJump)
	rl.UnloadSound(myGame.Player.FxHit)
	rl.UnloadSound(myGame.FxScore)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
