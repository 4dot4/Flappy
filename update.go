package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func jump(player *goppy) {

	player.SpeedY = 0
	player.SpeedY -= 6

}
func animePlayer(Player *goppy) {
	Player.FrameCounter++
	if Player.FrameCounter >= 60/Player.FrameSpeed {
		Player.FrameCounter = 0
		Player.CurrentFrame++
		if Player.CurrentFrame > 2 {
			Player.CurrentFrame = 0
		}
		Player.SourceRec.X = float32(Player.CurrentFrame) * Player.SourceRec.Width
	}
}
func fisica(Game *Game) {
	if Game.Player.CircleCol.Origin.Y+Game.Player.CircleCol.Radios >= Game.Foreground.RecDest.Y {
		Game.Player.SpeedY = 0
	} else {
		Game.Player.SpeedY += Gravity
	}
	if Game.Player.SpeedY > 10 {
		Game.Player.SpeedY = 10
	}
	if Game.Player.Rotation >= 85 {
		Gravity = 0.3
	} else {
		Gravity = 0.2
	}
	if Game.Player.SpeedY > 5 {
		if Game.Player.Rotation <= 90 {
			Game.Player.Rotation += 7
		}
	} else {
		if Game.Player.Rotation >= -40 {
			Game.Player.Rotation -= 10
		}
	}
	Game.Player.DestRec.Y += Game.Player.SpeedY
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
		jump(&Game.Player)
		Game.UpTube.DestRec.Y = float32(rl.GetRandomValue(-400, -150))
		Game.DownTube.DestRec.Y = Game.UpTube.DestRec.Y + Game.UpTube.DestRec.Height + 130

	}
}
func update(Game *Game) {
	Game.Foreground.ScrollF -= 2
	if Game.Foreground.ScrollF <= -Game.Foreground.RecDest.Width {
		Game.Foreground.ScrollF = 0
	}
	Game.Player.CircleCol.Origin = rl.Vector2{
		X: Game.Player.DestRec.X - 6,
		Y: Game.Player.DestRec.Y + 2,
	}

	animePlayer(&Game.Player)
	fisica(Game)
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
