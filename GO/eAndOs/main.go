package main

import "fmt"

func main() {
	tl := []int{1, 2, 3, 4, 5}

	for _, v := range tl {
		if v%2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else {
			fmt.Printf("%d is odd\n", v)
		}
	}
}
