package main

import (
	"fmt"
	"time"
	"unsafe"
	"strconv"
	"math/rand"
)

type Rect struct{
	x int
	y int
}

// Works for both Rect and Rect pointer
func (r Rect)String() string {
	return fmt.Sprintf("Rect x: %d, y: %d", r.x, r.y)
}

// Only pointers works, non pointer struct var still use the reflection mechanism to print out
// If uncomment, there will be error "method redeclared: Rect.String"
// func (pr *Rect)String() string {
// 	return fmt.Sprintf("Rect Pointer x: %d, y: %d", pr.x, pr.y)
// }
func main(){

	fmt.Println(rand.Intn(100))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))

	var r = Rect{1, 2}

	fmt.Println(r)
	fmt.Println(&r)

	n := 5
	ch := make(chan string)

	fmt.Println(unsafe.Sizeof(ch))

	for i :=0; i<n; i++{
		go func(index int){
			time.Sleep(time.Second)
			ch <- "ping from routine: " + strconv.Itoa(index)
		}(i)
	}

	for i :=0; i<n; i++{
		fmt.Println(<-ch)
	}

	// var pingCh = make(chan string, 1)
	// var pongCh = make(chan string, 1)

	// ping(pingCh, "ping msg")

}

func ping(ping chan <- string, msg string){
	ping <- msg
}

func pong(ping <- chan string, pong chan <- string){
	msg := <- ping
	pong <- msg
}