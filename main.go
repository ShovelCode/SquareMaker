package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func generatePNG(n int) {
	img := image.NewRGBA(image.Rect(0, 0, n, n))

	col := color.RGBA{0, 255, 0, 255} // Green

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			img.Set(x, y, col)
		}
	}

	file, err := os.Create(fmt.Sprintf("rectangle_%d.png", n))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	png.Encode(file, img)
}

func main() {
	generatePNG(100)
}
