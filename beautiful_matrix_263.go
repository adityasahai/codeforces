// https://codeforces.com/problemset/problem/263/A

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var numRows int = 5

func main() {
	var row, col int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < numRows; i++ {
		scanner.Scan()
		line := scanner.Text()
		if strings.Contains(line, "1") {
			row = i
			for j, char := range strings.Split(line, " ") {
				if char == "1" {
					col = j
				}
			}
		}
	}

	moves := math.Abs(2-float64(row)) + math.Abs(2-float64(col))
	fmt.Printf("%d", int(moves))
}
