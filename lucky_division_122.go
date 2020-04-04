// https://codeforces.com/problemset/problem/122/A
package main

import (
	"fmt"
	"strconv"
)

func checkNumLucky(number int) bool {
	numberString := strconv.Itoa(number)
	for _, r := range numberString {
		if !(r == '4' || r == '7') {
			return false
		}
	}
	return true
}

func checkNumLuckyProd(input int, luckyNumbers []int) {
	for _, ln := range luckyNumbers {
		if input%ln == 0 {
			fmt.Printf("YES")
			return
		}
	}
	fmt.Printf("NO")
}

func main() {
	luckyNumbers := []int{}
	for i := 1; i <= 1000; i++ {
		if checkNumLucky(i) {
			luckyNumbers = append(luckyNumbers, i)
		}
	}

	var input int
	fmt.Scanf("%d\n", &input)
	checkNumLuckyProd(input, luckyNumbers)
}
