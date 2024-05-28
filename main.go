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
