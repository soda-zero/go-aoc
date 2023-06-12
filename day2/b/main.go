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
	//  Paper 2
	//  Rock 1
	//  Scissors 3

	// Win  = 6
	// Draw = 3
	// Lose = 0

	// X = Lose
	// Y = Draw
	// Z = Win
	rock := 1
	paper := 2
	scissors := 3
	x := 0
	y := 3
	z := 6
	file, _ := os.Open("../day2.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	var score int
	scores := map[string]int{
		"A Y": rock + y,
		"A X": scissors + x,
		"A Z": paper + z,
		"B Y": paper + y,
		"B X": rock + x,
		"B Z": scissors + z,
		"C Y": scissors + y,
		"C X": paper + x,
		"C Z": rock + z,
	}
	for sc.Scan() {
		score += scores[sc.Text()]
	}
	fmt.Println(score)
}
