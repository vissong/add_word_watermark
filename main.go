package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

var fontface font.Face

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <image_directory> <watermark_text>")
		return
	}

	imageDir := os.Args[1]
	watermarkText := os.Args[2]

	files, err := filepath.Glob(imageDir + "/*.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fontface, err = gg.LoadFontFace("./65.ttf", 50)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		addWatermark(file, watermarkText)
	}
}

const (
	AddHeight = 0
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
