package main

import (
	"fmt"
	"os"
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

	cnt := make(map[int]int)
	for _, num := range b {
		cnt[num]++
	}

	ans := 0
	for _, num := range a {
		ans +=  num * cnt[num]
	}


	fmt.Println(ans)
}