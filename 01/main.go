package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1:", countIncreasing(getData()))
}

func partTwo() {
	var sums = []int{}
	var current = []int{}

	for _, val := range getData() {
		// add the new value to this next window
		current = append(current, val)
		if len(current) > 3 {
			current = current[1:4]
		}

		// if the window is full
		if len(current) == 3 {
			// add to the list of sums
			sums = append(sums, sum(current))
		}
	}

	// report the answer!
	fmt.Println("Part 2:", countIncreasing(sums))
}

// takes a list of integers, and tells how many of the
// values are higher than their previous value
func countIncreasing(in []int) (out int) {
	var increasing int
	var prev int
	for _, val := range in {
		if val > prev && prev != 0 {
			increasing++
		}
		prev = val
	}
	return increasing
}

// sum the integers in a list
func sum(in []int) (out int) {
	var s int
	for _, i := range in {
		s += i
	}
	return s
}

// getData will read our puzzles input data from a file
func getData() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	// scanner will read the file by lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data = []int{}

	for scanner.Scan() {
		// convert each line to an integer
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, val)
	}
	return data
}
