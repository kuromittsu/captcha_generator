package main

import (
	"math/rand"

	"github.com/fogleman/gg"
)

func noiseCircle(dc *gg.Context, rng *rand.Rand, loop int) {
	for i := 0; i < loop; i++ {
		// dc.SetRGBA(rng.Float64(), rng.Float64(), rng.Float64(), rng.Float64()*0.3)
		dc.SetRGBA(
			rng.Float64()*0.5, // 0.0–0.5
			rng.Float64()*0.5,
			rng.Float64()*0.5,
			0.2+rng.Float64()*0.2, // alpha 0.2–0.4
		)
		x := rng.Float64() * float64(width)
		y := rng.Float64() * float64(height)
		// size := rng.Float64()*8 + 3
		size := rng.Float64()*8 + 9
		dc.DrawCircle(x, y, size)
		dc.Fill()
	}
}
