package main

import "fmt"

// Cだと
// char* swap(char x,char y)
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
