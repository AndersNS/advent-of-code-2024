package main

import (
	"testing"
)

const (
	Part1ExampleExpected = 11
	Part2ExampleExpected = 31
)

func TestPart1(t *testing.T) {
	input := GetFile("example.txt")
	res, err := Part1(input)
	if err != nil {
		t.Fatal("Got err", err)
	}

	if res != Part1ExampleExpected {
		t.Errorf("Got %d, expected %d", res, Part1ExampleExpected)
	}
}

func TestPart2(t *testing.T) {
	input := GetFile("example.txt")
	res, err := Part2(input)
	if err != nil {
		t.Fatal("Got err", err)
	}

	if res != Part2ExampleExpected {
		t.Errorf("Got %d, expected %d", res, Part2ExampleExpected)
	}
}
