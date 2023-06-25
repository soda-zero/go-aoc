package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Open("./day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part_one(file)

	_, err = file.Seek(0, 0) // Rewind file pointer to the beginning
	if err != nil {
		log.Fatal(err)
	}

}

func part_one(file *os.File) {
}
