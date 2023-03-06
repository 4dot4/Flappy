package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	if Game.Player.CircleCol.Origin.Y+Game.Player.CircleCol.Radios >= ScreenHeight {
		Game.Player.SpeedY = -Game.Player.SpeedY
	} else {
		Game.Player.SpeedY += Gravity
	}
	if Game.Player.CircleCol.Origin.Y-Game.Player.CircleCol.Radios <= 0 {
		Game.Player.SpeedY = -Game.Player.SpeedY
	}
	if Game.Player.SpeedY > 0 {
		if Game.Player.Rotation <= 90 {
			Game.Player.Rotation += 3
		}
	} else {
		if Game.Player.Rotation >= -50 {
			Game.Player.Rotation -= 10
		}

	}
	Game.Player.DestRec.Y += Game.Player.SpeedY
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
		//Game.Player.FrameSpeed++
		//Game.Player.CircleCol.Origin.X += 1
		Game.Player.SpeedY = -Game.Player.SpeedY
	}
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
