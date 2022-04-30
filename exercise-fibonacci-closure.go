package main

import "fmt"

func fibonacci() func() int{
    i := 0
    j := 1

    return func() int {
        v := i
        i,j = j, i+j
        return v
    }
}

func main() {
    f := fibonacci()

    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
