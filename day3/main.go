package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("./day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part_one(file)

	_, err = file.Seek(0, 0) // Rewind file pointer to the beginning
	if err != nil {
		log.Fatal(err)
	}

	part_two(file)
}

func part_one(file *os.File) {
	sc := bufio.NewScanner(file)
	var sum int

	for sc.Scan() {
		items := make(map[rune]bool)
		for _, left := range sc.Text()[:len(sc.Text())/2] {
			items[left] = true
		}
		for _, right := range sc.Text()[len(sc.Text())/2:] {
			if items[right] {
				sum += calculatePriority(right)
				break
			}
		}
	}
	fmt.Println("Part one: ", sum)
}
func part_two(file *os.File) {
	sc := bufio.NewScanner(file)
	var sum int

	for sc.Scan() {
		itemsFirst := createSetItems(sc.Text())
		sc.Scan()
		itemsSecond := createSetItems(sc.Text())
		sc.Scan()
		itemsThird := createSetItems(sc.Text())

		for itemsFirstElf := range itemsFirst {
			if itemsSecond[itemsFirstElf] && itemsThird[itemsFirstElf] {
				sum += calculatePriority(itemsFirstElf)
				break
			}
		}

	}
    fmt.Println("Part two: ",sum)
}

func createSetItems(items string) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}
func hasItemInAll(items []map[rune]bool, item rune) bool {
	for _, itemSet := range items {
		if !itemSet[item] {
			return false
		}
	}
	return true
}
func calculatePriority(item rune) int {
	priority := int(unicode.ToLower(item) - 96)
	if unicode.IsUpper(item) {
		priority += 26
	}
	return priority
}
