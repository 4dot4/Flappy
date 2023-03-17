package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
		Radios: 20,
	}
	Game.Foreground.RecSource = rl.Rectangle{
		X:      292,
		Y:      0,
		Width:  167,
		Height: 56,
	}
	Game.Foreground.RecDest = rl.Rectangle{
		X:      0,
		Y:      570,
		Width:  ScreenWidth,
		Height: 200,
	}
	Game.UpTube = Tube{
		Source:  rl.Rectangle{X: 55, Y: 323, Width: 27, Height: 161},
		DestRec: rl.Rectangle{X: 300, Y: float32(rl.GetRandomValue(-400, -150)), Width: 27 * 3, Height: 3 * 161},
	}
	Game.DownTube = Tube{
		Source: rl.Rectangle{X: 83, Y: 320, Width: 27, Height: 161},
		DestRec: rl.Rectangle{
			X: Game.UpTube.DestRec.X,
			Y: Game.UpTube.DestRec.Y + Game.UpTube.DestRec.Height + 130, Width: 27 * 3, Height: 161 * 3},
	}
	Game.Player.FrameSpeed = 8
	Game.Player.SpeedY = 5
	Game.Over = false
	Game.Pause = false
	Game.Score = 0

}
func main() {

	var myGame Game

	rl.InitWindow(ScreenWidth, ScreenHeight, "flappy")
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
