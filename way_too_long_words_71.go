// https://codeforces.com/problemset/problem/71/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func abbvWord(word string) {
	if len(word) <= 10 {
		fmt.Println(word)
	} else {
		fmt.Printf("%c%d%c\n", word[0], len(word)-2, word[len(word)-1])
	}
}

func main() {
	var lines int
	fmt.Scanf("%d\n", &lines)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < lines; i++ {
		scanner.Scan()
		abbvWord(scanner.Text())
	}
}
