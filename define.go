package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen int

const (
	Menu Screen = iota
	Logo
	Scale float32 = 3

	ScreenWidth  = 400
	ScreenHeight = 700
)

var Gravity float32

type Circle struct {
	Origin rl.Vector2
	Radios float32
}
type goppy struct {
	DestRec      rl.Rectangle
	SourceRec    rl.Rectangle
	CircleCol    Circle
	FxJump       rl.Sound
	FxHit        rl.Sound
	Origin       rl.Vector2
	Rotation     float32
	CurrentFrame int
	FrameSpeed   int
	FrameCounter int
	SpeedY       float32
}
type Tube struct {
	Source  rl.Rectangle
	DestRec rl.Rectangle
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

	FxOver  rl.Music
	Player  goppy
	TubePos [200][2]Tube

	Foreground    Foreground
	HighScore     int32
	CurrentScreen Screen
}
