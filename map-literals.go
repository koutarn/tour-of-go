package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"bell labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		        37.42202, -122.08408,
	},}

var familiy = map[string]int {
    "Kouta":30,
    "Tomoko":33,
    "Mei":2,
}

func main(){
    fmt.Println(m)
    fmt.Println(familiy)
}
