package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Row struct {
	r     []int
	g     []int
	b     []int
	chars []rune
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Welcome to my command line cellular automata generator!")
	fmt.Println("Please enter the number of characters per line: ")
	var width int
	_, err := fmt.Scanf("%d", &width)
	if err != nil {
		log.Fatal("Invalid input")
	}
	prev := generateFirstRow(width)

	for {
		printRow(prev)
		prev = generateRow(prev)
	}

}

// generateFirstRow generates the first row of the cellular automata
func generateFirstRow(width int) Row {
	// Create a slice of ints for the RGB values and characters
	currentR := make([]int, width)
	currentG := make([]int, width)
	currentB := make([]int, width)
	currentChar := make([]int, width)

	// Randomly creating the first row
	for i := 0; i < width; i++ {
		currentR[i] = rand.Intn(256)
		currentG[i] = rand.Intn(256)
		currentB[i] = rand.Intn(256)
		// Chars MUST be in [32, 127)
		currentChar[i] = rand.Intn(94) + 33
	}

	return Row{
		r: currentR,
		g: currentG,
		b: currentB,
	}
}

// generateRow generates the next row based on the previous row
func generateRow(prev Row) Row {
	// Create a slice of ints for the RGB values and characters
	currentR := make([]int, len(prev.r))
	currentG := make([]int, len(prev.g))
	currentB := make([]int, len(prev.b))
	currentChar := make([]int, len(prev.chars)) // ASCII values for the characters

	// Randomly creating the first row
	for i := 0; i < len(prev.r); i++ {

		// FIXME
		currentR[i] = rand.Intn(256)
		currentG[i] = rand.Intn(256)
		currentB[i] = rand.Intn(256)
		// Chars MUST be in [32, 127)
		currentChar[i] = rand.Intn(95) + 32
	}

	return Row{
		r: currentR,
		g: currentG,
		b: currentB,
	}
}

// generateRValue generates the R value for the current cell based on the left, center, and right cells above it
func generateChar(left int, center int, right int) int {
	// Propagating spaces outwards
	if left == 32 || right == 32 {
		return 32
	}
	result := 0

	if left <= 79 {
		result += left
	}
	if center <= 79 {
		result += center
	}
	if right <= 79 {
		result += right
	}

	if result >= 127 {
		result = (result % 127) + 32
	}

	return result
}

// generateRValue generates the R value for the current cell based on the left, center, and right cells above it
func generateRValue(left int, center int, right int) int {

	return 0
}

// generateGValue generates the G value for the current cell based on the left, center, and right cells above it
func generateGValue(left int, center int, right int) int {

	return 0
}

// generateBValue generates the R value for the current cell based on the left, center, and right cells above it
func generateBValue(left int, center int, right int) int {

	return 0
}

// printRow prints the row to the terminal with the correct colors
func printRow(row Row) {
	for i := 0; i < len(row.r); i++ {
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", row.r[i], row.g[i], row.b[i], row.chars[i])
	}
	fmt.Println()
}
