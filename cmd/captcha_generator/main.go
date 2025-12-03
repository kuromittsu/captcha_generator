package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image/png"
	"math/rand"
	"os"
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

	// write text
	if fontPath == "" {
		scale := float64(height) / 20
		dc.Scale(scale, scale)
		dc.DrawStringAnchored(val,
			(float64(width)/2)/scale,
			(float64(height)/2)/scale,
			0.5, 0.5,
		)
		dc.Identity()
	} else {
		// scale := float64(height) / 20
		// dc.Scale(scale, scale)
		dc.DrawStringAnchored(val,
			float64(width)/2,
			float64(height)/2,
			0.5, 0.5,
		)
		dc.Identity()
	}

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

func outputBase64(dc *gg.Context) string {
	var buf bytes.Buffer
	png.Encode(&buf, dc.Image())
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func outputPng(dc *gg.Context, filename string) string {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, dc.Image())
	return filename
}
