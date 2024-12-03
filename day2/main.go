package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func main() {
	input := GetFile("input.txt")

	res, err := Part1(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1: ", res)

	res, err = Part2(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2: ", res)
}

func CheckLine(report []int) bool {
	dir := "u"
	for i := 0; i < len(report)-1; i += 1 {
		prev := report[i]
		level := report[i+1]

		diff := prev - level

		if diff < 0 {
			if dir == "a" {
				return false
			}
			if dir == "u" {
				dir = "d"
			}
		}
		if diff > 0 {
			if dir == "d" {
				return false
			}
			if dir == "u" {
				dir = "a"
			}
		}

		absDiff := math.Abs(float64(diff))

		if absDiff == 0 || absDiff > 3 {
			return false
		}
	}

	return true
}

func Part1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		columns, err := toInts(strings.Fields(line))
		if err != nil {
			return 0, err
		}

		valid := CheckLine(columns)
		if valid {
			sum++
		}
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		columns, err := toInts(strings.Fields(line))
		if err != nil {
			return 0, err
		}

		valid := CheckLine(columns)
		if valid {
			sum++
			continue
		}

		for i := range len(columns) {
			if CheckLine(removeAtIndex(columns, i)) {
				sum++
				break
			}
		}
	}

	return sum, nil
}

func removeAtIndex(slice []int, index int) []int {
	new := make([]int, 0, len(slice)-1)
	new = append(new, slice[:index]...)
	new = append(new, slice[index+1:]...)

	return new
}

func toInts(stringsArray []string) ([]int, error) {
	intArray := make([]int, len(stringsArray))
	for i, str := range stringsArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray[i] = num
	}
	return intArray, nil
}
