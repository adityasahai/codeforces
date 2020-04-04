package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(people string) {
	for _, p := range strings.Split(people, " ") {
		if p == "1" {
			fmt.Printf("hard")
			return
		}
	}
	fmt.Printf("easy")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	people := scanner.Text()
	solution(people)
}
