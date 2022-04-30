package main

import "fmt"


func add(x,y int) int {
    return x + y
}

func swap(x,y int) (int,int) {
    return y,x
}

func reverse(a,b,c,d string)(string,string,string,string){
    return d,c,b,a
}

func mul(sum int)(ret1,ret2 int){
    ret1 = sum / 2
    ret2 = sum / 2
    return
}

func main(){

    fmt.Println(add(10,20))
    fmt.Println(swap(10,20))
    fmt.Println(mul(100))
    fmt.Println(reverse("Hi","my","name","is"))

}
