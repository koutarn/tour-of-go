package main

import "fmt"

//Cだと
// int add(int x,int y)
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
