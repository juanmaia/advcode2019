package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/juanmaia/advcode2019/day2/intcode"
)

func main() {
	originalInput, err := SplitNumbers(os.Stdin, ",")
	if err != nil {
		panic(err)
	}
	i, j := 0, 0
Loop:
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			input := make([]int, len(originalInput))
			copy(input, originalInput)
			input[1] = i
			input[2] = j
			output, err := intcode.Run(input)
			if err != nil {
				panic(err)
			}
			fmt.Printf("i=%d j=%d output=%d\n", i, j, output[0])
			if output[0] == 19690720 {
				fmt.Printf("Breaking i=%d j=%d\n", i, j)
				break Loop
			}
		}
	}
	total := 100*i + j
	fmt.Println(total)
}

func SplitNumbers(r io.Reader, splitStr string) ([]int, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var result []int
	for _, str := range strings.Split(string(s), splitStr) {
		str = strings.TrimSuffix(str, "\n")
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}
