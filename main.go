package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen int

const (
	Menu Screen = iota
	Logo
	Scale float32 = 3
)

type Circle struct {
	Origin rl.Vector2
	Radios float32
}
type goppy struct {
	DestRec      rl.Rectangle
	FrameRec     rl.Rectangle
	CircleCol    Circle
	Rotation     float32
	CurrentFrame int
	FrameSpeed   int
	FrameCounter int
}
type tube struct {
	Texture rl.Texture2D
	Color   rl.Color
	Active  bool
}
type Game struct {
	Pause         bool
	Over          bool
	Score         int
	SpriteSheet   rl.Texture2D
	FxJump        rl.Sound
	FxOver        rl.Music
	Player        goppy
	Tube          tube
	HighScore     int
	CurrentScreen Screen
}

func update(Game *Game) {
	Game.Player.FrameCounter++
	if Game.Player.FrameCounter >= 60/Game.Player.FrameSpeed {
		Game.Player.FrameCounter = 0
		Game.Player.CurrentFrame++
		if Game.Player.CurrentFrame > 2 {
			Game.Player.CurrentFrame = 0
		}
		Game.Player.FrameRec.X = float32(Game.Player.CurrentFrame) * Game.Player.FrameRec.Width
	}
	Game.Player.Rotation += 1
	if rl.IsKeyPressed(rl.KeyD) {
		Game.Player.FrameSpeed++
	}
	if rl.IsKeyPressed(rl.KeyA) {
		Game.Player.FrameSpeed--
	}
	if rl.IsKeyDown(rl.KeyDown) {
		Game.Player.DestRec.Y += 5
	}
	if rl.IsKeyDown(rl.KeyUp) {
		Game.Player.DestRec.Y -= 5
	}
	if rl.IsKeyDown(rl.KeyRight) {
		Game.Player.DestRec.X += 5
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		Game.Player.DestRec.X -= 5
	}

}
func draw(Game *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	//background
	rl.DrawTexturePro(Game.SpriteSheet,
		rl.Rectangle{X: 0, Y: 0, Width: 143, Height: 255},         //RecSource
		rl.Rectangle{X: 0, Y: 0, Width: 143 * 4, Height: 255 * 4}, //Destiny
		rl.Vector2{X: 143 / 2, Y: 255 / 2}, 0, rl.White)           //Origin

	//floppy
	rl.DrawTexturePro(Game.SpriteSheet, Game.Player.FrameRec, Game.Player.DestRec, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	rl.DrawRectangleLines(Game.Player.DestRec.ToInt32().X+9, Game.Player.DestRec.ToInt32().Y, Game.Player.DestRec.ToInt32().Width-31, Game.Player.DestRec.ToInt32().Height, rl.Red)
	rl.DrawText(fmt.Sprintf("X : %2f\nY :%2f", Game.Player.DestRec.X, Game.Player.DestRec.Y), 10, 10, 20, rl.Red)
	rl.EndDrawing()
}
func (Game *Game) loadGame() {

	Game.SpriteSheet = rl.LoadTexture("./assets/SpriteSheet.png")

}
func (Game *Game) unloadGame() {

	rl.UnloadSound(Game.FxJump)
	rl.UnloadMusicStream(Game.FxOver)

}
func (Game *Game) initGame() {
	Game.Player = goppy{
		FrameRec: rl.Rectangle{
			X:      0,
			Y:      491,
			Width:  28,
			Height: 12,
		},
	}
	Game.Player.DestRec = rl.Rectangle{
		X:      140,
		Y:      360,
		Width:  Game.Player.FrameRec.Width * Scale,
		Height: Game.Player.FrameRec.Height * Scale,
	}

	Game.Player.FrameSpeed = 8

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
