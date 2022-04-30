package main

import "fmt"

type Vertex struct {
    Lat,Long float64
}

var m map[string]Vertex

var n map[string]int

func main(){
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])

    n = make(map[string]int)
    n["a"] = 1
    n["b"] = 2
    n["c"] = 3

    fmt.Printf("%d %d %d",n["a"],n["b"],n["c"])
}
