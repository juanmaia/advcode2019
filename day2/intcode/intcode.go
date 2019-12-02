package intcode

import (
	"errors"
	"fmt"
)

func Run(input []int) ([]int, error) {
	pos := 0

	// fmt.Println("Start:")
	// Print(input)
	for {
		if pos >= len(input) {
			return nil, errors.New("Not valid position")
		}
		opcode := input[pos]

		if opcode == 99 {
			// fmt.Println("end:")
			// Print(input)
			return input, nil
		}

		a, b, c := input[pos+1], input[pos+2], input[pos+3]
		if opcode == 1 {
			input[c] = input[a] + input[b]
		} else {
			input[c] = input[a] * input[b]
		}
		pos += 4
	}
	fmt.Println("End:")
	Print(input)

	return input, nil
}

func Print(input []int) {
	pos := 0
	fmt.Println("Program: ")

	for pos < len(input) {
		if input[pos] != 99 {
			fmt.Println(input[pos : pos+4])
			pos += 4
		} else {
			fmt.Println(input[pos])
			pos += 1
		}
	}
}
