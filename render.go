package main

type Render struct {
	Screen [][]Particle
}

func NewRender() *Render {
	//Create Initial Row Buffer
	grid := make([][]Particle, 30)
	for x := 0; x < len(grid); x++ {
		// For Every Row Creating a string Arr with 5 string
		grid[x] = make([]Particle, 30)
	}

	return &Render{
		Screen: grid,
	}
}
