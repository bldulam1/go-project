package main

import "fmt"

func main() {
	emails := map[string]string{"Bob": "bob@gmail.com",
		"Sharon": "sharon@gmail.com"}

	delete(emails, "Bob")

	fmt.Println(emails)
}
