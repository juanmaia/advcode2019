package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type reduceFunc func(int, int, int) int
type reduceFloatFunc func(float64, float64, int) float64

func main() {
	total := puzzle2()
	fmt.Println(total)
}

func puzzle1() int {
	total, err := reduceInt(os.Stdin, func(total, n, i int) int {
		return total + calculateMass(n)
	}, 0)
	if err != nil {
		panic(err)
	}
	return total
}

func puzzle2() float64 {
	total, err := reduceFloat(os.Stdin, func(total, n float64, i int) float64 {
		fmt.Println(n)
		for n > -1 {
			n = float64(calculateMass(int(n)))
			fmt.Printf("%f ", n)
			if n >= 0 {
				total += n
			}
		}
		fmt.Printf("Total = %f\n", total)
		return total
	}, 0)
	if err != nil {
		panic(err)
	}
	return total
}

func calculateMass(n int) int {
	n = n / 3
	return n - 2
}

func reduceInt(r io.Reader, f reduceFunc, initialValue int) (int, error) {
	total := initialValue
	i := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}
		total = f(total, n, i)
		i++
	}

	return total, nil
}

func reduceFloat(r io.Reader, f reduceFloatFunc, initialValue float64) (float64, error) {
	total := initialValue
	i := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			return 0, err
		}
		total = f(total, n, i)
		i++
	}

	return total, nil
}
