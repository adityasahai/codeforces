// https://codeforces.com/problemset/problem/339/A
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var problem string
	fmt.Scanf("%v\n", &problem)
	numbers := strings.Split(problem, "+")
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Printf("%v", strings.Join(numbers, "+"))
}
