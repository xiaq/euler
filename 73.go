package main

import "fmt"

func gcd(a, b int) int {
	for b > 0 {
		a = a % b
		a, b = b, a
	}
	return a
}

const n = 12000

func main() {
	count := 0
	for i := 1; i <= n; i++ {
		for j := i*2 + 1; j <= n && j < i*3; j++ {
			if gcd(i, j) == 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
