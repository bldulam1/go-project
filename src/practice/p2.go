package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	for i := 0; i < 100; i++ {
		if i%3 < 1 && i%5 < 1 {
			fmt.Println("fizzbuzz")
		} else if i%3 < 1 {
			fmt.Println("fizz")
		} else if i%5 < 1 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
