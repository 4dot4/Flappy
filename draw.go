package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(Game *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	//background
	rl.DrawTexturePro(
		Game.SpriteSheet, //texture
		rl.Rectangle{X: 0, Y: 0, Width: 143, Height: 255},             //RecSource
		rl.Rectangle{X: 0, Y: 0, Width: 143 * 3.5, Height: 255 * 3.5}, //Destiny
		rl.Vector2{X: 143 / 2, Y: 255 / 2}, 0, rl.White)               //Origin

	//floppy
	rl.DrawTexturePro(
		Game.SpriteSheet,      //Texture
		Game.Player.SourceRec, //SourceRec
		Game.Player.DestRec,   //Destiny
		Game.Player.Origin,    //Origin
		Game.Player.Rotation,  //Rotation
		rl.White)
	//uptube

	for i := 0; i < len(Game.TubePos); i++ {
		for d := 0; d < 2; d++ {
			rl.DrawTexturePro(
				Game.SpriteSheet,
				Game.TubePos[i][d].Source,
				Game.TubePos[i][d].DestRec,
				rl.Vector2{X: 0, Y: 0}, 0, rl.RayWhite)
		}
	}
	//Foreground
	rl.DrawTexturePro(
		Game.SpriteSheet,          //texture
		Game.Foreground.RecSource, //RecSource
		rl.Rectangle{
			X:      Game.Foreground.ScrollF,
			Y:      Game.Foreground.RecDest.Y,
			Width:  Game.Foreground.RecDest.Width,
			Height: Game.Foreground.RecDest.Height,
		}, //DestRec
		rl.Vector2{X: 0, Y: 0},
		0, rl.White)
	rl.DrawTexturePro(
		Game.SpriteSheet,          //texture
		Game.Foreground.RecSource, //RecSource
		rl.Rectangle{
			X:      Game.Foreground.RecDest.Width + Game.Foreground.ScrollF,
			Y:      Game.Foreground.RecDest.Y,
			Width:  Game.Foreground.RecDest.Width,
			Height: Game.Foreground.RecDest.Height,
		}, //DestRec
		rl.Vector2{X: 0, Y: 0},
		0, rl.White)

	rl.DrawCircleLines(int32(Game.Player.CircleCol.Origin.X), int32(Game.Player.CircleCol.Origin.Y), Game.Player.CircleCol.Radios, rl.Red) //physics
	rl.DrawCircle(int32(Game.Player.CircleCol.Origin.X), int32(Game.Player.CircleCol.Origin.Y), 5, rl.Red)
	rl.DrawText(fmt.Sprintf("Score %d", Game.Score), 10, 10, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("Velocity: %f \nGravity:%f", Game.Player.SpeedY, Gravity), 10, 60, 20, rl.Red)
	rl.EndDrawing()
}
