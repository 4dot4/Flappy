package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var bigger int = -1
var Start bool = false
var ySpeed float32 = -1
var tube int = -1
var OverTest = false

func tubesCol(Game *Game) {

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
				if tube != i {
					rl.PlaySound(Game.FxScore)
					Game.Score++

					tube = i
				}

			}

			if Game.TubePos[i][d].DestRec.X+Game.TubePos[i][d].DestRec.Width < 0 {
				if bigger == -1 {
					Game.TubePos[i][d].DestRec.X = Game.TubePos[len(Game.TubePos)-1][0].DestRec.X +
						Game.TubePos[i][d].DestRec.Width + Xspace

					bigger = 0
				} else {
					Game.TubePos[i][d].DestRec.X = Game.TubePos[bigger][0].DestRec.X +
						Game.TubePos[i][d].DestRec.Width + Xspace
					Game.TubePos[i][0].DestRec.Y = float32(rl.GetRandomValue(-400, -150))

					bigger = i
				}

				Game.TubePos[i][1].DestRec.Y = Game.TubePos[i][0].DestRec.Y + Game.TubePos[i][0].DestRec.Height + Yspace
			}

		}

	}
}
func restart(Game *Game) {
	Game.Over = false
	OverTest = false
	Game.Player.Rotation = 0
	Start = false
	Game.Score = 0
	tube = -1
	bigger = -1
	startTubes(&Game.TubePos, &Game.Player)
}
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

func playerMov(Game *Game) {
	if Start {
		if Game.Player.CircleCol.Origin.Y+Game.Player.CircleCol.Radios >= Game.Foreground.RecDest.Y {

			Game.Player.SpeedY = 0
			if !OverTest {
				rl.PlaySound(Game.FxOver)
				OverTest = true
			}
			Game.Over = true

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
			if Game.Player.Rotation >= -30 {
				Game.Player.Rotation -= 10
			}
		}
		Game.Player.DestRec.Y += Game.Player.SpeedY
	} else {
		Game.Player.DestRec.Y += ySpeed
		if Game.Player.DestRec.Y > 330 {
			ySpeed = -ySpeed
		}
		if Game.Player.DestRec.Y < 390 {
			ySpeed = -ySpeed
		}
	}

}
func fisica(Game *Game) {

	tubesCol(Game)
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
		Start = true
		jump(&Game.Player)

	}
}
func update(Game *Game) {

	animePlayer(&Game.Player)

	playerMov(Game)

	Game.Player.CircleCol.Origin = rl.Vector2{
		X: Game.Player.DestRec.X - 6,
		Y: Game.Player.DestRec.Y + 2,
	}
	if !Game.Over {

		Game.Foreground.ScrollF -= 3
		if Game.Foreground.ScrollF <= -Game.Foreground.RecDest.Width {
			Game.Foreground.ScrollF = 0
		}
		if Start {
			for i := 0; i < len(Game.TubePos); i++ {
				Game.TubePos[i][0].DestRec.X -= 3
				Game.TubePos[i][1].DestRec.X = Game.TubePos[i][0].DestRec.X
			}

		}
		fisica(Game)
	}

	if rl.IsKeyDown(rl.KeyR) {
		restart(Game)
	}
	if rl.IsKeyPressed(rl.KeyF) {

		if ShowFPS {
			ShowFPS = false
		} else {
			ShowFPS = true
		}

	}
}
