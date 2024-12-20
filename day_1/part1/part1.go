package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	input, _ := os.Open("input.txt")
	output, _ := os.Create("output.txt")
	os.Stdin = input
	os.Stdout = output

	a, b := []int{}, []int{}
	for i := 0; i < 1000; i++{
		var x, y int
		fmt.Scanf("%d   %d", &x, &y)
		a = append(a, x)
		b = append(b, y)
	}

	sort.Ints(a)
	sort.Ints(b)
	ans := 0

	for i := 0; i < 1000; i++ {
		val := a[i] - b[i]
		if val < 0 {
			val *= -1
		}
		ans += val
	}

	fmt.Print(ans)

}