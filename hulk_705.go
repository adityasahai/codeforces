// https://codeforces.com/problemset/problem/705/A
package main

import (
	"fmt"
	"strings"
)

var love string = "I love"
var hate string = "I hate"
var end string = "it"
var conj string = "that"

func main() {
	var layers int
	fmt.Scanf("%d\n", &layers)

	response := make([]string, 0)

	for i:= 1; i <= layers; i++ {
		if i > 1 {
			response = append(response, conj)
		}
		if i % 2 == 1 {
			response = append(response, hate)
		} else {
			response = append(response,love)
		}
	}
	response = append(response, end)

	fmt.Printf("%+v\n", strings.Join(response," "))

}