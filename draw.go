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

	//Foreground
	rl.DrawTexturePro(
		Game.SpriteSheet, //texture
		rl.Rectangle{X: 292, Y: 0, Width: 167, Height: 56},          //RecSource
		rl.Rectangle{X: 0, Y: 600, Width: ScreenWidth, Height: 200}, //DestRec
		rl.Vector2{X: 0, Y: 0},
		0, rl.White)

	//floppy
	rl.DrawTexturePro(
		Game.SpriteSheet,      //Texture
		Game.Player.SourceRec, //SourceRec
		Game.Player.DestRec,   //Destiny
		Game.Player.Origin,    //Origin
		Game.Player.Rotation,  //Rotation
		rl.White)
	rl.DrawCircleLines(int32(Game.Player.CircleCol.Origin.X), int32(Game.Player.CircleCol.Origin.Y), Game.Player.CircleCol.Radios, rl.Red) //physics
	rl.DrawText(fmt.Sprintf("X : %2f\nY :%2f", Game.Player.DestRec.X, Game.Player.DestRec.Y), 10, 10, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("Velocity: %f \nGravity:%f", Game.Player.SpeedY, Gravity), 10, 60, 20, rl.Red)

	rl.EndDrawing()
}
