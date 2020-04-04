// https://codeforces.com/problemset/problem/236/A
package main

import (
	"fmt"
)

func is_girl(name string) bool {
	chars := map[rune]bool{}
	for _, char := range name {
		chars[char] = true
	}
	return len(chars)%2 == 0
}

func main() {
	var name string
	fmt.Scanf("%v\n", &name)
	if is_girl(name) {
		fmt.Printf("CHAT WITH HER!")
	} else {
		fmt.Printf("IGNORE HIM!")
	}
}
