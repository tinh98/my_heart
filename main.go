package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"math/rand"
)

const (
	width    = 500
	height   = 500
	fontPath = "../Library/Fonts/Atomic-Love.ttf"
)

// Define some parameters for the animation
const (
	numParticles  = 200
	maxAge        = 10
	particleSpeed = 3
	minRadius     = 2
	maxRadius     = 5
)

func main() {
	initView()
	loadView()
	//time.Sleep(500 * time.Second)
}

func initView() {
	dc := gg.NewContext(width, height)

	if err := dc.LoadFontFace(fontPath, 48); err != nil {
		panic(err)
	}

	particles := make([]particle, numParticles)
	for i := 0; i < numParticles; i++ {
		// Initialize each particle with a random position, velocity, and age
		particles[i].x = randFloat(width)
		particles[i].y = randFloat(height)
		particles[i].vx = randFloat(-particleSpeed, particleSpeed)
		particles[i].vy = randFloat(-particleSpeed, particleSpeed)
		particles[i].age = randInt(maxAge)
		particles[i].radius = randFloat(minRadius, maxRadius)
		particles[i].color = color.RGBA{255, 180, 0, 255}
	}
	zome := 0.2
	max := 5.0
	scale := 1.0 // or any other scaling factor
	min := scale

	// Main animation loop
	for v := 0; ; v++ {
		// Clear the canvas
		dc.SetColor(color.Black)
		dc.Clear()

		// Update and draw each particle
		for i := 0; i < numParticles; i++ {
			particles[i].update()
			particles[i].draw(dc)
		}

		// Draw text
		dc.SetColor(color.RGBA{0, 120, 30, 255})
		drawText(v, width, dc)

		// Draw a heart shape in the center of the canvas
		dc.SetColor(color.RGBA{255, 0, 0, 255})
		//dc.DrawPath(heartShape(width/2, height/2, 100))
		//for t := 0.0; t <= 2.0*math.Pi; t += 0.01 {
		//	x := 16 * math.Pow(math.Sin(t), 3)
		//	y := 13*math.Cos(t) - 5*math.Cos(2*t) - 2*math.Cos(3*t) - math.Cos(4*t)
		//	dc.LineTo(width/2+x, height/2+y)
		//}

		if math.Abs(scale) >= max || math.Abs(scale) < min {
			zome *= -1
		}

		for t := 0.0; t <= 2.0*math.Pi; t += 0.01 {
			x := scale * 44 * math.Sin(t) * math.Sin(t) * math.Sin(t)
			y := scale * (-1 * (32*math.Cos(t) - 12*math.Cos(2*t) - 5*math.Cos(3*t) - 2*math.Cos(4*t)))
			dc.LineTo(width/2+x, height/2+y)
		}
		scale += zome
		dc.Fill()

		// Save the canvas as a PNG file
		dc.SavePNG(fmt.Sprintf("frames/heart_animation_%d.png", v))

		// Pause for a short time to control the frame rate
		//time.Sleep(time.Millisecond * 50)

		if v > 300 {
			break
		}
	}
}

type particle struct {
	x, y, vx, vy float64
	age          int
	radius       float64
	color        color.RGBA
}

func (p *particle) update() {
	p.x += p.vx
	p.y += p.vy
	p.radius += 1
	p.age++
	if p.age > maxAge {
		// If the particle has reached its maximum age, reset it with a new position and age
		p.x = randFloat(width)
		p.y = randFloat(height)
		p.age = 0
		p.radius = randFloat(minRadius, maxRadius)
	}
}

func (p *particle) draw(dc *gg.Context) {
	alpha := 1.0 - float64(p.age)/float64(maxAge)
	dc.SetColor(color.RGBA{p.color.R, p.color.G, p.color.B, uint8(alpha * 255)})
	dc.DrawEllipse(p.x, p.y, p.radius, p.radius)
	dc.Fill()
}

//func heartShape(x, y, size float64) gg.Path {
//	path := gg.NewPath()
//	path.MoveTo(x, y)
//	for angle := 0.0; angle < math.Pi*2; angle += 0.01 {
//		rx := size * 16 * math.Pow(math.Sin(angle), 3)
//		ry := -size * (13*math.Cos(angle) - 5*math.Cos(2*angle) - 2*math.Cos(3*angle) - math.Cos(4*angle))
//		path.LineTo(x+rx, y+ry)
//	}
//	path.Close()
//	return path
//}

func randFloat(args ...float64) float64 {
	if len(args) == 1 {
		return rand.Float64() * args[0]
	} else if len(args) == 2 {
		return args[0] + rand.Float64()*(args[1]-args[0])
	} else {
		panic("Invalid number of arguments")
	}
}

//func randFloat(max float64) float64 {
//	return rand.Float64() * max
//}

func randInt(max int) int {
	return rand.Intn(max)
}

func drawText(v, width int, dc *gg.Context) {
	w := float64((v*2 + 1) % width)
	dc.DrawString("Happy Valentine's Day!", w, 60)
}
