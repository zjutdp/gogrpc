package main

import (
  "fmt"
  "hash/crc32"
)

func main(){
  h := crc32.NewIEEE()
  fmt.Println(h)
  fmt.Println(h.Sum32())
}
