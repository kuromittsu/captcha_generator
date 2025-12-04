package main

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/fogleman/gg"
)

func outputBase64(dc *gg.Context) string {
	var buf bytes.Buffer

	// png.Encode(&buf, dc.Image())

	// base64 get from JPG format
	jpeg.Encode(&buf, dc.Image(), &jpeg.Options{Quality: 80})

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
