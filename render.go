package main

import "fmt"

type Render struct {
	Screen [][]string
}

func NewRender() *Render {

	//Create Initial Row Buffer
	grid := make([][]string, 30)
	for x := 0; x < len(grid); x++ {
		// For Every Row Creating a string Arr with 5 string
		grid[x] = make([]string, 30)
	}
	// return the Render the object
	return &Render{
		Screen: grid,
	}
}

// Render the board
func (r *Render) RenderBoard() {
	// fmt.Println(len(r.Screen))
	for i := 0; i < len(r.Screen); i++ {
		fmt.Println(r.Screen[i])
	}
}
