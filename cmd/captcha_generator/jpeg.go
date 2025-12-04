package main

import (
	"image/jpeg"
	"os"

	"github.com/fogleman/gg"
)

func outputJpg(dc *gg.Context, filename string) string {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jpeg.Encode(file, dc.Image(), &jpeg.Options{
		Quality: 80,
	})
	return filename
}
