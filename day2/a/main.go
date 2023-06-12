package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Opponent
	// A: Rock
	// B: Paper
	// C: Scissors

	//Myself
	// Y: Paper 2
	// X: Rock 1
	// Z: Scissors 3

	// Win  = 6
	// Draw = 3
	// Lose = 0
	file, _ := os.Open("../day2.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	var score int
	scores := map[string]int{
		"A Y": 8,
		"A X": 4,
		"A Z": 3,
		"B Y": 5,
		"B X": 1,
		"B Z": 9,
		"C Y": 2,
		"C X": 7,
		"C Z": 6,
	}
	for sc.Scan() {
		score += scores[sc.Text()]
	}
    fmt.Println(score)
}
