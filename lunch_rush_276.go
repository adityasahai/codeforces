// https://codeforces.com/problemset/problem/276/A
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"math"
)

func getTwoNum(text string) (int64, int64) {
	inputs := strings.Split(text, " ")
	input1, _ := strconv.ParseInt(inputs[0], 10, 32)
	input2, _ := strconv.ParseInt(inputs[1], 10, 32)
	return input1, input2
}
 
func calculateJoy(fi, ti, k int64) (joy int64) {
	if ti > k {
		joy = fi - (ti - k)
	} else {
		joy = fi
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numRest, time := getTwoNum(scanner.Text())
	fun := int64(math.MinInt64)
	var i int64
	for i = 0; i < numRest; i++ {
		scanner.Scan()
		fi, ti := getTwoNum(scanner.Text())
		j := calculateJoy(fi, ti, time)
		if j > fun {
			fun = j
		}
	}
	fmt.Printf("%d", fun)
}