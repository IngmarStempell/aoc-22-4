package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 22 No 4")
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Finding total encloses areas")
	sum := 0
	for _, line := range readLine() {
		sum += hasTotalOverlap(split(line))
	}
	fmt.Printf("Found %d total enclosed areas \n", sum)
}

func split(input string) (string, string) {
	splits := strings.Split(input, ",")
	return splits[0], splits[1]
}

func hasTotalOverlap(left string, right string) int {
	startLeft := getStart(left)
	endLeft := getEnd(left)

	startRight := getStart(right)
	endRight := getEnd(right)

	if (startLeft <= startRight) && (endLeft >= endRight) {
		return 1
	}
	if startRight <= startLeft && endRight >= endLeft {
		return 1
	}
	return 0
}

func hasPartialOverlap(left string, right string) int {
	startLeft := getStart(left)
	endLeft := getEnd(left)

	startRight := getStart(right)
	endRight := getEnd(right)

	if (startLeft <= startRight) && (endLeft >= startRight) {
		return 1
	}
	if startRight <= startLeft && endRight >= startLeft {
		return 1
	}
	return 0
}

func getEnd(input string) int {
	splits := strings.Split(input, "-")
	v, _ := strconv.Atoi(splits[1])
	return v
}

func getStart(input string) int {
	splits := strings.Split(input, "-")
	v, _ := strconv.Atoi(splits[0])
	return v
}

func partTwo() {
	fmt.Println("Finding partial encloses areas")
	sum := 0
	for _, line := range readLine() {
		sum += hasPartialOverlap(split(line))
	}
	fmt.Printf("Found %d partial enclosed areas \n", sum)
}

func readLine() []string {
	fileHandle, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can not open input file")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(fileHandle)
	lines := []string{}
	line, err := reader.ReadString('\n')
	for line != "" {
		if err != nil {
			fmt.Println("Cant read from file")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		lines = append(lines, strings.Trim(line, "\n"))
		line, err = reader.ReadString('\n')
	}
	return lines
}
