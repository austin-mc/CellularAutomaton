/*
Austin Christiansen
Cellular Automata Generator
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var animation gif.GIF

const (
	// Initial setup
	width      = 600
	height     = 600
	squareSize = width / 100
	frameCount = 101 // Total number of frames
	frameDelay = 5   // 5 ms delay between frames
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

var showGrid bool
var ruleset int

func main() {
	var input string
	fmt.Println("Welcome to my cellular automata generator!")
	fmt.Println("Would you like to show the grid lines? y/n (default: n)")
	fmt.Scanln(&input)
	if input == "y" {
		showGrid = true
	} else {
		showGrid = false
	}
	fmt.Println("Please select a cellular automata ruleset (1-5):")
	fmt.Println("1 - 3 are more standard cellular automata rulesets")
	fmt.Println("4 - 5 are experimental rulesets using bit manipulation")
	fmt.Scanln(&input)
	inputVal, err := strconv.Atoi(input)
	if err != nil || inputVal < 1 || inputVal > 5 {
		fmt.Println("Invalid input, exiting...")
		os.Exit(1)
	}
	ruleset = inputVal

	fmt.Println("Generating animation...")
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

	out, err := os.Create("out.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = gif.EncodeAll(out, &animation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Animation saved to out.gif")
}

// Following cellular automata rules to update the grid
// Each index must be an int in range [0, 5)
func updateGrid(grid [][]int, row int) [][]int {

	if row == 0 {
		// First row is all white except the center cell
		// Which will be given a random color
		middle := 50
		grid[row][middle] = rand.Intn(4) + 1
	} else {
		prevRow := grid[row-1]
		currentRow := grid[row]

		var left int
		var center int
		var right int

		for i := 0; i < 100; i++ {
			// getting adjacent 3 cells from the previous row
			// If the cell is on the edge, wrap around to the other side
			if i == 0 {
				left = prevRow[99]
			} else {
				left = prevRow[i-1]
			}
			if i == 99 {
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
// Has 5 rulesets to choose from
func generateCell(left int, center int, right int) int {
	if ruleset == 1 {
		return (left + center + right) % 5
	}
	if ruleset == 2 {
		leftSum := left + center
		rightSum := right + center
		return (leftSum + rightSum) % 5
	}
	if ruleset == 3 {
		if left == 0 && right == 0 && center == 0 {
			return 0
		}
		if left == 0 && right == 0 {
			return center
		}
		if left == 0 && center == 0 {
			return right
		}
		if right == 0 && center == 0 {
			return left
		}
		if left == 0 {
			return (center * right) % 5
		}
		if right == 0 {
			return (center * left) % 5
		}
		return (left * right) % 5
	}
	if ruleset == 4 {
		// Bit manipulation
		return ((right << left) ^ center) % 5
	}
	// ruleset == 5
	// More bit manipulation
	return ((left << right) ^ center) % 5

}

func drawNextFrame(width int, height int, squareSize int, grid [][]int) *image.Paletted {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			switch grid[y][x] {
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
	if x != 0 && showGrid {
		startX++
	}
	startY := y * squareSize
	if y != 0 && showGrid {
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
