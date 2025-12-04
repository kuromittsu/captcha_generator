package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/basicfont"
)

var width, height int
var val, out, fontPath string

func init() {
	flag.StringVar(&val, "val", "", "value")

	flag.IntVar(&width, "w", 160, "image width")
	flag.IntVar(&height, "h", 60, "image height")

	flag.StringVar(&out, "out", "base64", "base64 atau png")

	flag.StringVar(&fontPath, "fontpath", "", "path to ttf font")

	flag.Parse()
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	dc := gg.NewContext(width, height)

	bgR, bgG, bgB := setBackgound(dc, rng)

	if fontPath != "" {
		if err := dc.LoadFontFace(fontPath, float64(height/2)); err != nil {
			dc.SetFontFace(basicfont.Face7x13)
		}
	} else {
		dc.SetFontFace(basicfont.Face7x13)
	}

	// a bit rotation to make seem like captcha
	// angle := float64(rng.Intn(35) - 25) // between -25 to +10 deg
	angle := float64(rng.Intn(31) - 15)
	dc.RotateAbout(gg.Radians(angle), float64(width/2), float64(height/2))

	lum := 0.299*bgR + 0.587*bgG + 0.114*bgB
	if lum > 0.5 {
		dc.SetRGB(0.0, 0.0, 0.0) // light background → dark text
	} else {
		dc.SetRGB(1.0, 1.0, 1.0) // dark background → light text
	}

	// draw text
	drawCharsWithRotation(dc, val, rng)

	// circle noise
	noiseCircle(dc, rng, 10) // before 40

	// line noise
	noiseLine(dc, rng, 5) // before 20

	switch out {
	case "base64":
		output := outputBase64(dc)
		fmt.Print(output) // output to stdout

	case "png":
		output := outputPng(dc, fmt.Sprintf("captcha_%s.png", val))
		fmt.Printf("successfully generate : %s", output)

	case "jpg":
		output := outputJpg(dc, fmt.Sprintf("captcha_%s.jpg", val))
		fmt.Printf("successfully generate : %s", output)

	default:
		fmt.Print("")
	}

}

func randColor(rng *rand.Rand) float64 {
	return 0.2 + rng.Float64()*0.6 // clamp antara 0.2–0.8
}

func setBackgound(dc *gg.Context, rng *rand.Rand) (float64, float64, float64) {
	bgR := randColor(rng)
	bgG := randColor(rng)
	bgB := randColor(rng)

	dc.SetRGB(bgR, bgG, bgB)
	dc.Clear()

	return bgR, bgG, bgB
}

func drawCharsWithRotation(dc *gg.Context, text string, rng *rand.Rand) {
	n := len(text)
	charWidth := float64(width) / float64(n)

	for i, ch := range text {
		dc.Push() // save context

		x := charWidth*float64(i) + charWidth/2
		y := float64(height) / 2

		// random rotation per character
		angle := float64(rng.Intn(41) - 20) // -20 deg to +20 deg
		dc.RotateAbout(gg.Radians(angle), x, y)

		// slightly randomize the Y position to make it more natural
		y += float64(rng.Intn(7) - 3) // -3 px to +3 px

		dc.DrawStringAnchored(string(ch), x, y, 0.5, 0.5)

		dc.Pop() // restore context
	}
}
