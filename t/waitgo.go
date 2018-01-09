package main

import (
    "fmt"
    "time"
)
func main(){
    var wg sync.waitgroup

    for i := 0; i < ; i++ {
        fmt.Printf("%d ", i)
    }
    wg.Add(1)

    go func() {
    for i := 0; i < ; i++ {
        fmt.Printf("%d ", i)
        }
    defer wg.Done()
    }
    wg.Wait()

}
