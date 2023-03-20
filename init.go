package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func startTubes(TubePos *[5][2]Tube) {

	for i := 0; i < len(TubePos); i++ {
		if i == 0 {
			TubePos[i][0] = Tube{
				Source:  rl.Rectangle{X: 55, Y: 323, Width: 27, Height: 161},
				DestRec: rl.Rectangle{X: 850, Y: float32(rl.GetRandomValue(-400, -150)), Width: 27 * 3, Height: 3 * 161},
			}
		} else {
			TubePos[i][0] = Tube{
				Source: rl.Rectangle{X: 55, Y: 323, Width: 27, Height: 161},
				DestRec: rl.Rectangle{X: TubePos[i-1][0].DestRec.X + TubePos[i-1][0].DestRec.Width + Xspace,
					Y:      float32(rl.GetRandomValue(-400, -150)),
					Width:  27 * 3,
					Height: 3 * 161},
			}
		}
		TubePos[i][1] = Tube{
			Source: rl.Rectangle{X: 83, Y: 320, Width: 27, Height: 161},
			DestRec: rl.Rectangle{
				X:      TubePos[i][0].DestRec.X,
				Y:      TubePos[i][0].DestRec.Y + TubePos[i][0].DestRec.Height + Yspace,
				Width:  27 * 3,
				Height: 161 * 3},
		}
	}

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
		Radios: 17,
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
	startTubes(&Game.TubePos)
	Game.Player.FrameSpeed = 8
	Game.Player.SpeedY = 5
	Game.Over = false
	Game.Pause = false
	Game.Score = 0
}
