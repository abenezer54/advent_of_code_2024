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
		
		for removed := range a {
			pos, neg := 0, 0
			safe := true
			for i := 0; i < len(a) - 1; i++{
				if i == removed {
					continue
				}

				if i + 1 == removed {
					if i + 2 < len(a) {
						abs := a[i] - a[i + 2]
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
				} else {
					abs := a[i] - a[i + 1]
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

			}

			if pos > 0 && neg > 0 {
				safe = false
			}
			if safe {
				ans++
				break
			}
		}

	}

	fmt.Println(ans)
}