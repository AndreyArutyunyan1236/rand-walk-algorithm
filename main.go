package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Left = iota   // 0 
	Right        // 1 
	Up          // 2 
	Down       // 3
	
)

func rand_walk(startingPositionX int32, startingPositionY int32, color rl.Color) {
	escDontPressed := true
	fps := 0 
	for escDontPressed {
		if rl.IsKeyPressed(rl.KeyEscape) || fps == 800 { escDontPressed = false }	
		
		PositionX, PositionY := startingPositionX, startingPositionY

		rand_num := rand.Intn(4) 

		if rand_num == 0 { startingPositionX = PositionX - 20 }
		if rand_num == 1 { startingPositionX = PositionX + 20 }
		if rand_num == 2 { startingPositionY = PositionY + 20 }
		if rand_num == 3 { startingPositionY = PositionY - 20 }

		rl.DrawLine(startingPositionX, startingPositionY, PositionX, PositionY, color)
		fps++
	}
}

func main() {
	rl.InitWindow(1300, 600, "Random Walk")
	
	rl.SetTargetFPS(2)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		
		rl.ClearBackground(rl.Black)
		// logic is here

		rand_walk(400, 400, rl.Green)
		rand_walk(400, 400, rl.Blue)
		rand_walk(400, 400, rl.Red)
		rand_walk(400, 400, rl.Yellow)

		rand_walk(400, 400, rl.NewColor(128, 0, 128, 255))
		rand_walk(400, 400, rl.NewColor(255, 69, 0, 255))

		rl.EndDrawing()
	}
}

