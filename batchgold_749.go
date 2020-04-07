// https://codeforces.com/problemset/problem/749/A
package main

import (
	"fmt"
	"strings"
)

func addtwos(ans []string, times int) []string {
	for i := 0; i < times; i++ {
		ans = append(ans, "2")
	}
	return ans
}

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	ans := []string{}

	remtwo := num / 2
	divtwo := num % 2 == 0
	if divtwo {
		ans = addtwos(ans, remtwo)
	} else {
		ans = addtwos(ans, remtwo-1)
		ans = append(ans, "3")
	}
	fmt.Printf("%d\n", len(ans))
	fmt.Printf("%v", strings.Join(ans, " "))
}