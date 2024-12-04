package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	input, _ := os.Open("./../input.txt")
	output, _ := os.Create("./../output.txt")
	os.Stdin = input
	os.Stdout = output

	reader := bufio.NewScanner(os.Stdin)
	ans := 0
	for reader.Scan() {
		line := strings.Split(reader.Text(), " ")
		a := []int{}
		for i := range line {
			num, _ := strconv.Atoi(line[i])
			a = append(a, num)
		}

		pos, neg := 0, 0
		safe := true
		for i := range a {
			if i == 0 {
				continue
			}
			abs := a[i] - a[i - 1]
			if abs < 0 {
				neg++
				abs *= -1
			} else {
				pos++
			}
			if !(1 <= abs  && abs <= 3){
				safe = false
				break
			}
			
		}
		if pos > 0 && neg > 0 {
			safe = false
		}
		if safe {
			ans++
		}


	}
	fmt.Println(ans)
}