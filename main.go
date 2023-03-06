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
	SourceRec    rl.Rectangle
	CircleCol    Circle
	Origin       rl.Vector2
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

	Game.Player.CircleCol.Origin = rl.Vector2{
		X: Game.Player.DestRec.X - 6,
		Y: Game.Player.DestRec.Y,
	}
	Game.Player.FrameCounter++
	if Game.Player.FrameCounter >= 60/Game.Player.FrameSpeed {
		Game.Player.FrameCounter = 0
		Game.Player.CurrentFrame++
		if Game.Player.CurrentFrame > 2 {
			Game.Player.CurrentFrame = 0
		}
		Game.Player.SourceRec.X = float32(Game.Player.CurrentFrame) * Game.Player.SourceRec.Width
	}
	Game.Player.Rotation += 1
	//if rl.IsKeyPressed(rl.KeyD) {
	//Game.Player.FrameSpeed++
	//	Game.Player.CircleCol.Origin.X += 1
	//}
	// if rl.IsKeyPressed(rl.KeyA) {
	// 	//Game.Player.FrameSpeed--
	// 	Game.Player.CircleCol.Origin.X -= 1
	// }
	// if rl.IsKeyPressed(rl.KeyW) {
	// 	//Game.Player.FrameSpeed++
	// 	Game.Player.CircleCol.Origin.Y -= 1
	// }
	// if rl.IsKeyPressed(rl.KeyS) {
	// 	//Game.Player.FrameSpeed--
	// 	Game.Player.CircleCol.Origin.Y += 1
	// }
	// if rl.IsKeyPressed(rl.KeyF) {
	// 	//Game.Player.FrameSpeed--
	// 	Game.Player.CircleCol.Radios += 1
	// }
	// if rl.IsKeyPressed(rl.KeyG) {
	// 	Game.Player.CircleCol.Radios -= 1
	// }
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
	rl.DrawTexturePro(
		Game.SpriteSheet, //texture
		rl.Rectangle{X: 0, Y: 0, Width: 143, Height: 255},         //RecSource
		rl.Rectangle{X: 0, Y: 0, Width: 143 * 4, Height: 255 * 4}, //Destiny
		rl.Vector2{X: 143 / 2, Y: 255 / 2}, 0, rl.White)           //Origin

	//floppy
	rl.DrawTexturePro(
		Game.SpriteSheet,      //Texture
		Game.Player.SourceRec, //SourceRec
		Game.Player.DestRec,   //Destiny
		Game.Player.Origin,    //Origin
		Game.Player.Rotation,  //Rotation
		rl.White)
	rl.DrawCircleLines(int32(Game.Player.CircleCol.Origin.X), int32(Game.Player.CircleCol.Origin.Y), Game.Player.CircleCol.Radios, rl.Red)
	rl.DrawText(fmt.Sprintf("X : %2f\nY :%2f", Game.Player.DestRec.X, Game.Player.DestRec.Y), 10, 10, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("CircleX : %2f\nCircleY :%2f, Radios: %f", Game.Player.CircleCol.Origin.X, Game.Player.CircleCol.Origin.Y, Game.Player.CircleCol.Radios), 10, 60, 20, rl.Red)

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
		SourceRec: rl.Rectangle{
			X:      0,
			Y:      491,
			Width:  28,
			Height: 12,
		},
	}
	Game.Player.DestRec = rl.Rectangle{
		X:      140,
		Y:      360,
		Width:  Game.Player.SourceRec.Width * Scale,
		Height: Game.Player.SourceRec.Height * Scale,
	}
	Game.Player.Origin = rl.Vector2{
		X: Game.Player.DestRec.Width / 2,
		Y: Game.Player.DestRec.Height / 2,
	}
	Game.Player.CircleCol = Circle{
		Origin: rl.Vector2{
			X: Game.Player.DestRec.X,
			Y: Game.Player.DestRec.Y,
		},
		Radios: 21,
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
