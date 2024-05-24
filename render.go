package main

import (
	"fmt"
)

type Render struct {
	Board []string
}

func NewRender() *Render {
	//Create Initial Row Buffer
	grid := make([]string, 10)
	for x := 0; x < len(grid); x++ {
	}
	// return the Render the object
	return &Render{
		Board: grid,
	}
}

func addParticles(particle *Particle, board []string) {
	for i, _ := range board {
		board[i] = particle.ascii
	}
}

// Render the board
func (r *Render) RenderBoard() {
	particle := createParticle()
	addParticles(particle, r.Board)
	fmt.Println(r.Board)
}
