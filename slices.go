package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

    //1から3までの要素を指定した事になる
    var s []int = primes[1:4]
    fmt.Println(s)

    var d []int
    d = primes[0:6]
    fmt.Println(d)
}
