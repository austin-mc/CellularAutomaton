/*
Austin Christiansen
Cellular Automata Generator
*/

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math/rand"
	"time"
)

var animation gif.GIF

const (
	// Initial setup
	width      = 600
	height     = 600
	squareSize = width / 100
	frameCount = height + 1 // Total number of frames is height + 1 since the first grid is the initial grid
	frameDelay = 5          // 5 ms delay between frames
)

/*
Global Colors
*/

// Base Colors
var white = color.RGBA{249, 249, 249, 255}
var black = color.RGBA{0, 0, 0, 255}

// Cell Colors (https://www.color-hex.com/color-palette/5452)
var red = color.RGBA{217, 83, 79, 255}
var lightBlue = color.RGBA{91, 192, 222, 255}
var green = color.RGBA{92, 184, 92, 255}
var darkBlue = color.RGBA{66, 139, 202, 255}

var myPallette = color.Palette{
	white,
	black,
	red,
	lightBlue,
	green,
	darkBlue,
}

func main() {
	animation = gif.GIF{LoopCount: frameCount}
	// Random used only to generate color of the first cell
	rand.Seed(time.Now().Unix())

	// Grid values will be 0 initially indicating white color
	grid := make([][]int, 100)
	for i := range grid {
		grid[i] = make([]int, 100)
	}

	var img *image.Paletted
	for i := 0; i < frameCount; i++ {
		if i != 0 {
			grid = updateGrid(grid, i-1)
		}
		img = drawNextFrame(width, height, squareSize, grid)
		appendImage(img)
	}

}

// Following cellular automata rules to update the grid
// Each index must be an int in range [0, 5)
func updateGrid(grid [][]int, row int) [][]int {

	if row == 0 {
		// First row is all white except the center cell
		// Which will be given a random color
		middle := width / 2
		grid[row][middle] = rand.Intn(4) + 1
	} else {
		prevRow := grid[row-1]
		currentRow := grid[row]

		var left int
		var center int
		var right int

		for i := 0; i < width; i++ {
			// getting adjacent 3 cells from the previous row
			// If the cell is on the edge, wrap around to the other side
			if i == 0 {
				left = prevRow[width-1]
			} else {
				left = prevRow[i-1]
			}
			if i == width-1 {
				right = prevRow[0]
			} else {
				right = prevRow[i+1]
			}
			center = prevRow[i]

			currentRow[i] = generateCell(left, center, right)
		}
	}
	return grid
}

// Generate the color for the cell given the 3 previous colors
func generateCell(left int, center int, right int) int {
	if left == 0 && center == 0 && right == 0 {
		return 0
	}

}

// // Draws the initial grid for the automata
// func drawGrid(width int, height int, squareSize int) *image.Paletted {
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)
// 	for x := 0; x < 100; x++ {
// 		for y := 0; y < 100; y++ {
// 			// Draw the squares to make the grid
// 			drawSquare(x, y, squareSize, white, img)
// 		}
// 	}

// 	// Convert to image.Paletted to be used in gif
// 	palettedImage := image.NewPaletted(img.Bounds(), myPallette)
// 	draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Src)
// 	return palettedImage
// }

func drawNextFrame(width int, height int, squareSize int, grid [][]int) *image.Paletted {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			switch grid[x][y] {
			case 0:
				drawSquare(x, y, squareSize, white, img)
			case 1:
				drawSquare(x, y, squareSize, red, img)
			case 2:
				drawSquare(x, y, squareSize, lightBlue, img)
			case 3:
				drawSquare(x, y, squareSize, green, img)
			case 4:
				drawSquare(x, y, squareSize, darkBlue, img)
			}
		}
	}

	// Convert to image.Paletted to be used in gif
	palettedImage := image.NewPaletted(img.Bounds(), myPallette)
	draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Src)
	return palettedImage
}

// Draws a square on image m at the specified x and y coordinates
func drawSquare(x int, y int, squareSize int, color color.RGBA, m *image.RGBA) {
	startX := x * squareSize
	if x != 0 {
		startX++
	}
	startY := y * squareSize
	if y != 0 {
		startY++
	}
	endX := (x * squareSize) + squareSize
	endY := (y * squareSize) + squareSize
	square := image.Rect(startX, startY, endX, endY)
	draw.Draw(m, square, &image.Uniform{color}, image.Point{}, draw.Src)
}

// Appends the image to the animation variable
func appendImage(img *image.Paletted) {
	animation.Image = append(animation.Image, img)
	animation.Delay = append(animation.Delay, frameDelay)
}
