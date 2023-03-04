package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type goppy struct {
	Coords rl.Vector2
}
type tube struct {
	Texture rl.Texture2D
	Color   rl.Color
	Active  bool
}
type Game struct {
	Pause       bool
	Over        bool
	Score       int
	SpriteSheet rl.Texture2D
	FxJump      rl.Sound
	FxOver      rl.Music
	Player      goppy
	Tube        tube
	HighScore   int
}

func update(Game *Game) {

}
func draw(Game *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.GetColor(0x052c46ff))
	rl.DrawTextureEx(Game.SpriteSheet, rl.Vector2{X: 0, Y: 0}, 0, 3, rl.White)
	rl.EndDrawing()
}
func (Game *Game) loadGame() {

	Game.FxJump = rl.LoadSound("./sounds/jump.mp3")
	Game.SpriteSheet = rl.LoadTexture("./assets/SpriteSheet.png")
	Game.FxOver = rl.LoadMusicStream("./sounds/gameOver.mp3")

}
func (Game *Game) unloadGame() {

	rl.UnloadSound(Game.FxJump)
	rl.UnloadMusicStream(Game.FxOver)

}
func (Game *Game) initGame() {

	Game.HighScore = 0
	Game.Over = false
	Game.Pause = false
	Game.Score = 0

}
func main() {

	var myGame Game

	rl.InitWindow(400, 700, "flappy")
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
