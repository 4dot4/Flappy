package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type goppy struct {
	coords rl.Vector2
}
type tubes struct {
	rec    rl.Rectangle
	color  rl.Color
	active bool
}
type game struct {
	pause      bool
	over       bool
	score      int
	background rl.Texture2D
	foreground rl.Texture2D
	midground  rl.Texture2D
	fxJump     rl.Sound
	fxOver     rl.Music
	player     goppy
	highScore  int
}

func update(game *game) {

}
func draw(game *game) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(game.background, 0, 0, rl.Green)

	rl.EndDrawing()
}
func (game *game) loadGame() {
	game.fxJump = rl.LoadSound("./sounds/jump.mp3")
	game.background = rl.LoadTexture("./assets/background.png")
	game.foreground = rl.LoadTexture("./assets/foreground.png")
	game.midground = rl.LoadTexture("./assets/midground.png")
	game.fxOver = rl.LoadMusicStream("./sounds/gameOver.mp3")

}
func (game *game) unloadGame() {
	rl.UnloadTexture(game.background)
	rl.UnloadTexture(game.foreground)
	rl.UnloadTexture(game.midground)
	rl.UnloadSound(game.fxJump)
	rl.UnloadMusicStream(game.fxOver)

}
func (game *game) initGame() {
	game.highScore = 0
	game.over = false
	game.pause = false
	game.score = 0

	game.player = goppy{rl.Vector2{X: 0, Y: 400}}
}
func main() {

	var myGame game

	rl.InitWindow(800, 450, "flappy")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	myGame.loadGame()
	myGame.initGame()

	for !rl.WindowShouldClose() {
		update(&myGame)
		draw(&myGame)
	}
	myGame.unloadGame()
	rl.CloseWindow()
}
