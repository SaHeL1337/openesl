package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi       = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile  = flag.String("fontfile", "Xolonium-Bold.ttf", "filename of the ttf font")
	imagePath = "images/"
)

func renderImage(item *Item) error {

	template, err := parseTemplate()
	if err != nil {
		return err
	}
	img := image.NewRGBA(image.Rect(0, 0, int(template.Width), int(template.Height)))

	background := color.RGBA{uint8(template.Background[0]), uint8(template.Background[1]), uint8(template.Background[2]), uint8(template.Background[3])}
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.Point{X: 0, Y: 0}, draw.Src)

	renderText(img, &template)

	path := path.Join(imagePath, strconv.Itoa(item.Id)+".png")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
	return nil
}

func renderText(img *image.RGBA, template *Template) {

	for _, text := range template.Text {
		drawTextToImage(img, text.X, text.Y, float64(text.Size), text.Content, color.RGBA{uint8(text.Color[0]), uint8(text.Color[1]), uint8(text.Color[2]), uint8(text.Color[3])})
	}

}

func drawTextToImage(img *image.RGBA, x, y int, fontsize float64, label string, color color.RGBA) {

	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	fontBytes, err := os.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(color),
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    fontsize,
			DPI:     *dpi,
			Hinting: font.HintingFull,
		}),
		Dot: point,
	}
	d.DrawString(label)
}
