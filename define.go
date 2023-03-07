package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen int

const (
	Menu Screen = iota
	Logo
	Scale        float32 = 3
	Gravity      float32 = 0.07
	ScreenWidth          = 400
	ScreenHeight         = 700
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
	SpeedY       float32
}
type tube struct {
	Texture rl.Texture2D
	Color   rl.Color
	Active  bool
}
type Foreground struct {
	RecSource rl.Rectangle
	RecDest   rl.Rectangle
	ScrollF   float32
}
type Game struct {
	Pause       bool
	Over        bool
	Score       int
	SpriteSheet rl.Texture2D
	FxJump      rl.Sound
	FxOver      rl.Music
	Player      goppy
	Tube        tube
	Foreground  Foreground
	HighScore   int

	CurrentScreen Screen
}
