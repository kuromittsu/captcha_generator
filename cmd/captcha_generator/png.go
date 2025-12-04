package main

import (
	"image/png"
	"os"

	"github.com/fogleman/gg"
)

func outputPng(dc *gg.Context, filename string) string {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, dc.Image())
	return filename
}
