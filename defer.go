# learn defer

package main

import (
   "fmt"
)

func main(){
   a()
   b()
   c := c()

   fmt.Println(c)
}

func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}

func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Println(i)
    }
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}
