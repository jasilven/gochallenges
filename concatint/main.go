package main

import (
	"fmt"
	"os"
	"strconv"
)

// concat given list of positive integers so that resulting integer is the
// largest integer possible
func main() {
	ints := make([]string, len(os.Args)-1)
	copy(ints, os.Args[1:])

	for i := 1; i < len(ints); i++ {
		for j := i; j < len(ints); j++ {
			next, _ := strconv.Atoi(ints[j] + ints[i-1])
			old, _ := strconv.Atoi(ints[i-1] + ints[j])

			if next > old {
				tmp := ints[j]
				ints[j] = ints[i-1]
				ints[i-1] = tmp
			}
		}
	}

	for _, s := range ints {
		fmt.Printf("%v", s)
	}
}
