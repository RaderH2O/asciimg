package main

import (
	"image"
	"image/color"

	"github.com/nfnt/resize"
)

func findAverageColor(colors ...color.Color) color.Color {
	average := color.RGBA{0, 0, 0, 255}
	averageValue := uint8(0)
	for _, col := range colors {
		r, g, b, _ := col.RGBA()
		// fmt.Printf("r - %v | g - %v | b - %v\n", r, g, b)
		red := (float64(r) / float64(65535)) * float64(255)
		green := (float64(g) / float64(65535)) * float64(255)
		blue := (float64(b) / float64(65535)) * float64(255)

		// averageValue += uint8(((float64(r) / float64(len(colors))) + (float64(g) / float64(len(colors))) + (float64(b) / float64(len(colors)))) / 3)
		averageValue += uint8(red / float64(len(colors)))
		averageValue += uint8(green / float64(len(colors)))
		averageValue += uint8(blue / float64(len(colors)))

		// NOTE: for colored images, not that useful
		// average.R += uint8(int(r) / len(colors))
		// average.G += uint8(int(g) / len(colors))
		// average.B += uint8(int(b) / len(colors))
	}

	average.R = averageValue
	average.G = averageValue
	average.B = averageValue

	return average
}

func handleImage(img image.Image, asciiWidth int) image.Image {
	// width := uint(58)
	width := uint(asciiWidth)
	height := uint(float32(width) / float32(img.Bounds().Max.X) * float32(img.Bounds().Max.Y))

	newImage := resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor)
	return newImage

	// topLeft := image.Point{0, 0}
	// bottomRight := image.Point{width, height}
	//
	// newImg := image.NewRGBA(image.Rectangle{topLeft, bottomRight})
	//
	// for i := range width {
	// 	for j := range height {
	// 		// average := findAverageColor(img.At(i*2, j*2), img.At(i*2+1, j*2), img.At(i*2, j*2+1), img.At(i*2+1, j*2+1))
	// 		// fmt.Println("I'm doing something!")
	// 		pixels := []color.Color{}
	// 		multiplier := int(img.Bounds().Max.X / width)
	// 		for counterX := range multiplier {
	// 			for counterY := range multiplier {
	// 				pixels = append(pixels, img.At(i*multiplier+counterX, j*multiplier+counterY))
	// 			}
	// 		}
	// 		average := findAverageColor(pixels...)
	// 		newImg.Set(i, j, average)
	// 	}
	// }
	//
	// // width := newImg.Bounds().Max.X
	// //
	// // if width < 100 {
	// // 	return *newImg
	// // }
	// // return handleImage(newImg)
	// return newImg
}
