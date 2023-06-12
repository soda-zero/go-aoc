package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var mostCals int
	var currentCals int

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())
		currentCals += snack
		if err != nil {
			if currentCals > mostCals {
				mostCals = currentCals
			}
			currentCals = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mostCals)
}
