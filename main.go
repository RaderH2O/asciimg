package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"strconv"
)

func main() {
	// pixelValues := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'."
	// pixelValues := "`.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"
	// pixelValues := ".'\":-+=0O%$##"
	pixelValues := ".:-=+*#%%"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the path to the image you want to use >>> ")
	scanner.Scan()
	filename := scanner.Text() // The image filename

	width := 0
	var err error = nil
	for {
		fmt.Print("Enter the width of the ASCII you want to create >>> ")
		scanner.Scan()
		width, err = strconv.Atoi(scanner.Text())
		if err != nil || width <= 0 {
			fmt.Println("Please enter a valid number for the width!")
			continue
		}
		break
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image: ", err)
		return
	}

	var f *os.File
	f, _ = os.Create("handled.png")
	defer f.Close()
	resizedImage := handleImage(img, width)
	png.Encode(f, resizedImage)

	output := ""
	for j := range resizedImage.Bounds().Max.Y {
		for i := range resizedImage.Bounds().Max.X {
			r, g, b, _ := resizedImage.At(i, j).RGBA()
			avg := int((r + g + b) / 3)

			output += string(pixelValues[int((float64(avg)/65764)*(float64(len(pixelValues)-1)))])
		}
		output += "\n"
	}

	fmt.Println(output)
	os.WriteFile("message.txt", []byte(output), 0666)

}
