package main

import (
	"fmt"
)

func getString() []int {
	retVal := make([]int, 2)
	return retVal
}

func getEntriesLen() int {
	var entriesLen int
	fmt.Scan(&entriesLen)
	return entriesLen
}

func main() {
	// var casesLen int
	var str int
	var str2 int
	// fmt.Scan(&casesLen)

	format := "%d\n"
	format = format + "%d "
	fmt.Scanf(format, &str, &str2)
	fmt.Println(str, str2)
}
