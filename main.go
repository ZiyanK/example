package main

import "fmt"

func Add(n1, n2 int) int {
	return n1 + n2
}

func main() {
	output := Add(1, 2)
	fmt.Println(output)
}
