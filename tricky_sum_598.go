// https://codeforces.com/problemset/problem/598/A
package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "math"
)

// func isPowerOfTwo(num int64) (float64, bool) {
//     pow := math.Log2(float64(num))
//     return pow, pow == math.Trunc(pow)
// }

func getSum(maxPow int64) int64 {
    return int64(math.Pow(2, float64(maxPow+1))) - 1
}

func calculateSum(count int64) int64 {
    var sum int64 = 0
    sum = int64(float64(count * (count + 1)) / 2) 
    pow := math.Log2(float64(count))
    maxPow := int64(math.Trunc(pow))
    return sum - 2 * getSum(maxPow)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    numInts, _ := strconv.ParseInt(scanner.Text(), 10, 32)
    var i int64
    for i = 0; i < numInts; i++ {
        scanner.Scan()
        count, _ := strconv.ParseInt(scanner.Text(), 10, 32)
        fmt.Printf("%d\n", calculateSum(count))
    }
}