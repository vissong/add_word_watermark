package main

import (
	"log"
	"path/filepath"

	"github.com/fogleman/gg"
)

func addWatermark(filename, watermark string) {
	im, err := gg.LoadImage(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file: %s, width: %d, height: %d", filename, im.Bounds().Dx(), im.Bounds().Dy())
	width := im.Bounds().Dx()
	height := im.Bounds().Dy() + AddHeight

	watermarkX := 30
	watermarkY := height - AddHeight/2 - 50

	dc := gg.NewContext(width, height)
	// background
	dc.SetHexColor("#FFFFFF")
	dc.Clear()
	// image
	dc.DrawImage(im, 0, 0)
	// word
	// dc.Clear()
	// if err := dc.LoadFontFace("./w7.ttf", 36); err != nil {
	// if err := dc.LoadFontFace("./65.ttf", 50); err != nil {
	// 	log.Fatal(err)
	// }
	dc.SetFontFace(fontface)
	dc.SetHexColor("#000000")
	dc.DrawStringAnchored(watermark, float64(watermarkX), float64(watermarkY), 0, 0.5)

	// save
	err = dc.SavePNG("./out/" + filepath.Base(filename))
	if err != nil {
		log.Fatal(err)
	}
}
