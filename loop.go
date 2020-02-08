package main

import "fmt"

func main() {
	i := 0

	fmt.Println("No", i)

	i++
	if i == 5 {
		break
	}
}
