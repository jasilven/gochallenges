package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// fill matrix with numbers 1..n in spiral pattern.
// matrix size (rows=cols) is read from program argument
func main() {
	size, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	m := getMatrix(int(size), int(size))

	x, y := 0, 0
	dx, dy := 1, 0

	for count := 1; count <= int(size*size); count++ {
		m[y][x] = count
		if (x+dx >= int(size)) || (x+dx < 0) || (y+dy >= int(size)) || (y+dy < 0) || (m[y+dy][x+dx] != 0) {
			tmp := dy
			dy = dx
			dx = -tmp
		}
		x = x + dx
		y = y + dy
	}
	printMatrix(m)
}

func getMatrix(r, c int) [][]int {
	m := make([][]int, r)
	for i := range m {
		m[i] = make([]int, c)
	}
	return m
}

func printMatrix(m [][]int) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			fmt.Printf("%3d ", m[y][x])
		}
		fmt.Println()
	}
}
