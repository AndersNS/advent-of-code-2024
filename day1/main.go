package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

func Part1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	var col1, col2 []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		num1, err1 := strconv.Atoi(columns[0])
		num2, err2 := strconv.Atoi(columns[1])

		if err1 != nil {
			return 0, err1
		}
		if err2 != nil {
			return 0, err2
		}

		if len(columns) >= 2 {
			col1 = append(col1, num1)
			col2 = append(col2, num2)
		}
	}

	sort.Ints(col1)
	sort.Ints(col2)

	sum := 0
	for i := range col1 {
		num1 := col1[i]
		num2 := col2[i]

		diff := math.Abs(float64(num2 - num1))

		sum += int(diff)
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	var col1, col2 []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		num1, err1 := strconv.Atoi(columns[0])
		num2, err2 := strconv.Atoi(columns[1])

		if err1 != nil {
			return 0, err1
		}
		if err2 != nil {
			return 0, err2
		}

		if len(columns) >= 2 {
			col1 = append(col1, num1)
			col2 = append(col2, num2)
		}
	}

	sort.Ints(col1)
	sort.Ints(col2)

	sum := 0
	for _, num := range col1 {
		count := count(col2, num)
		sum += (num * count)
	}

	return sum, nil
}

func count(arr []int, num int) int {
	count := 0
	for _, v := range arr {
		if v == num {
			count++
		}
	}
	return count
}
