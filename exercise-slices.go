package main

// import "golang.org/x/tour/pic"
// import "fmt"

func Pic(dx,dy int) [][]uint8{
    ret := make([][]uint8, dy)

    for i :=0; i < dy;i++ {
        ret[i] = make([]uint8,dx)
    }

    return ret
}

func main() {
    // pic.Show(Pic)
}
