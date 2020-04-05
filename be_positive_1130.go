// http://codeforces.com/problemset/problem/1130/A
package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"math"
	"strconv"
)

func main() {
	var inputs int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	line := scanner.Text()
	var zeros int = 0
	var positives int = 0
	var negatives int = 0
	for _, num := range strings.Split(line, " ") {
		if num == "0" {
			zeros++
		} else if strings.Count(num, "-") == 1 {
			negatives++
		} else {
			positives++
		}
	}
	c := math.Ceil(float64(inputs)/2)
	if float64(positives) >= c {
		fmt.Printf("1")
	} else if float64(negatives) >= c {
		fmt.Printf("-1")
	} else {
		fmt.Printf("0")
	}
}