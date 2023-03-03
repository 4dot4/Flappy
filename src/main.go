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
	pause         bool
	over          bool
	score         int
	background    rl.Texture2D
	foreground    rl.Texture2D
	midground     rl.Texture2D
	fxJump        rl.Sound
	fxOver        rl.Music
	player        goppy
	scrollingBack float32
	scrollingMid  float32
	scrollingFore float32
	highScore     int
}

func update(game *game) {
	game.scrollingBack -= 0.1
	game.scrollingMid -= 0.5
	game.scrollingFore -= 1

	if game.scrollingBack <= float32(game.background.Width)*2 {
		game.scrollingBack = 0
	}
	if game.scrollingMid <= -float32(game.midground.Width)*2 {
		game.scrollingMid = 0
	}
	if game.scrollingFore <= -float32(game.foreground.Width)*2 {
		game.scrollingFore = 0
	}

}
func draw(game *game) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x052c46ff))
	rl.DrawTextureEx(game.background, rl.Vector2{X: game.scrollingBack, Y: 20}, 0, 2.0, rl.White)
	rl.DrawTextureEx(game.background, rl.Vector2{X: float32(game.background.Width)*2 + game.scrollingBack, Y: 20}, 0, 2.0, rl.White)

	rl.DrawTextureEx(game.midground, rl.Vector2{X: game.scrollingMid, Y: 20}, 0, 2.0, rl.White)
	rl.DrawTextureEx(game.midground, rl.Vector2{X: float32(game.midground.Width)*2 + game.scrollingMid, Y: 20}, 0, 2.0, rl.White)

	rl.DrawTextureEx(game.foreground, rl.Vector2{X: game.scrollingFore, Y: 70}, 0, 2.0, rl.White)
	rl.DrawTextureEx(game.foreground, rl.Vector2{X: float32(game.foreground.Width)*2 + game.scrollingFore, Y: 70}, 0, 2.0, rl.White)
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
