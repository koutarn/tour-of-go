package main

import (
    // "golang.org/x/tour/wc"
    "fmt"
    "strings"
)


func WordCount(s string) map[string]int {

    sf := strings.Fields(s)
    fmt.Println(sf)
    
    ret := make(map[string]int)

    for i,_ := range sf {
        (ret[sf[i]])++
        fmt.Println(ret)
    }

    return ret
}

func main(){
    // wc.Test(WordCount)
    WordCount("This is a pen")
}
