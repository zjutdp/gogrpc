package main

import "fmt"

func main(){
  fmt.Printf("Hello, %s!", new(World))
}

type World struct{}

func (w *World) String() string{
  return "World"
}
