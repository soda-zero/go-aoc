package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentCals int
	var elfCalArray []int

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())
		currentCals += snack
		if err != nil {
			elfCalArray = append(elfCalArray, currentCals)
			currentCals = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elfCalArray)

	lastThree := elfCalArray[len(elfCalArray)-3:]
	var topThreeCaloriesTotal int

	for _, value := range lastThree {
		topThreeCaloriesTotal += value
	}

	fmt.Println(topThreeCaloriesTotal)
}
