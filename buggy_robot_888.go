// https://codeforces.com/problemset/problem/888/B
/* 
Consider the final cell after original path. It has some distance dx to x = 0 and dy to y = 0. That means the path included at least dx and dy in corresponding directions. Let's remove just these minimal numbers of moves.

Finally, the answer will be n - dx - dy, where (dx, dy) are distances from the final cell of the original path to (0, 0).

Overall complexity: O(n).
*/
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

func move(pos []int, move rune) {
	switch move {
	case 'U':
		pos[1]++
	case 'D':
		pos[1]--
	case 'L':
		pos[0]--
	case 'R':
		pos[0]++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numCommands, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	pos := []int{0, 0}
	scanner.Scan()
	cmds := scanner.Text()
	for _, cmd := range cmds {
		move(pos, cmd)
		// fmt.Printf("%d - current position - %d, %d\n", i+1, pos[0], pos[1])
	}
	fmt.Printf("%d", numCommands - int64(math.Abs(float64(pos[0]- 0))) - int64(math.Abs(float64(pos[1]-0))))
}