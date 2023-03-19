package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func jump(player *goppy) {

	rl.PlaySound(player.FxJump)
	player.SpeedY = 0
	player.SpeedY -= 5.5

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
func rewrite(TubePos [][]Tube) {

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

	for i := 0; i < len(Game.TubePos); i++ {
		for d := 0; d < 2; d++ {
			if Game.Player.CircleCol.Origin.X+Game.Player.CircleCol.Radios >= Game.TubePos[i][d].DestRec.X &&
				Game.Player.CircleCol.Origin.X-Game.Player.CircleCol.Radios <= Game.TubePos[i][d].DestRec.X+Game.TubePos[i][d].DestRec.Width &&
				Game.Player.CircleCol.Origin.Y+Game.Player.CircleCol.Radios >= Game.TubePos[i][d].DestRec.Y &&
				Game.Player.CircleCol.Origin.Y-Game.Player.CircleCol.Radios <= Game.TubePos[i][d].DestRec.Y+Game.TubePos[i][d].DestRec.Height {
				rl.PlaySound(Game.Player.FxHit)
				Game.Over = true
			}
			if Game.TubePos[i][d].DestRec.X+Game.TubePos[i][d].DestRec.Width < Game.Player.CircleCol.Origin.X+Game.Player.CircleCol.Radios {
				Game.Score = i + 1
			}

		}
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
		jump(&Game.Player)

	}
}
func update(Game *Game) {

	animePlayer(&Game.Player)
	if !Game.Over {

		Game.Foreground.ScrollF -= 2
		if Game.Foreground.ScrollF <= -Game.Foreground.RecDest.Width {
			Game.Foreground.ScrollF = 0
		}
		Game.Player.CircleCol.Origin = rl.Vector2{
			X: Game.Player.DestRec.X - 6,
			Y: Game.Player.DestRec.Y + 2,
		}

		for i := 0; i < len(Game.TubePos); i++ {
			Game.TubePos[i][0].DestRec.X -= 3
			Game.TubePos[i][1].DestRec.X = Game.TubePos[i][0].DestRec.X
		}
		fisica(Game)
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
