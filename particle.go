package main

import (
	r "math/rand"
	"time"
)

type Particle struct {
	ascii    string
	x        int
	y        int
	lifeTime int
}

var ascii = []string{"@", "}", "{"}

func NewParticle(ascii string, x, y int) *Particle {
	return &Particle{
		ascii:    ascii,
		x:        x,
		y:        y,
		lifeTime: time.Now().Minute(),
	}
}

func createParticle() *Particle {
	randX := r.Intn(10)
	randY := r.Intn(10)
	randAscii := r.Intn(3)
	particle := NewParticle(ascii[randAscii], randX, randY)
	return particle
}
