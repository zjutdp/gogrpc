package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice_dy := make([][]uint8, dy)

	for y := range slice_dy {
		slice_dy[y] = make([]uint8, dx)
		for x := range slice_dy[y] {
			v := (x*x + y*y) * 2
			//v := x*x + y*y
			//v := (x * x * y * y)
			slice_dy[y][x] = uint8(v)
		}
	}

	return slice_dy
}


func main() {
	pic.Show(Pic)

}

