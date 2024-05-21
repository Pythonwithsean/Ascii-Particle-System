package main

type Particle struct {
	data string
	x    int
	y    int
}

func NewParticle(data string, x, y int) *Particle {
	return &Particle{
		data: data,
		x:    x,
		y:    y,
	}
}

func (p *Particle) moveLeft() {}

func (p *Particle) moveRight() {}
