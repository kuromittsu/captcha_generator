package main

import (
	"math/rand"

	"github.com/fogleman/gg"
)

func noiseLine(dc *gg.Context, rng *rand.Rand, loop int) {
	for i := 0; i < loop; i++ {
		dc.SetRGBA(rng.Float64(), rng.Float64(), rng.Float64(), 0.8)
		dc.SetLineWidth(float64(rng.Intn(3) + 1))
		x1 := rng.Float64() * float64(width)
		y1 := rng.Float64() * float64(height)
		x2 := rng.Float64() * float64(width)
		y2 := rng.Float64() * float64(height)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
}
