package main

import "fmt"

type Vertex struct {
    X,Y,Z int
}

var (
    v1 = Vertex{1,2,3}
    v2 = Vertex{X:1}
    v3 = Vertex{}
    v4 = Vertex{Z:4,Y:2}
    p = &Vertex{1,2,3}
)

func main(){
    fmt.Println(v1,p,v2,v3,v4)
}
