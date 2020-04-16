package main

import ("fmt")

var a, x, y, p, sq, x1, y1 int

func main() {
	// you can change this values
	a = 5
	x = 0
	y = 0
	// now we start our calculations
	sq = a*a
	p = a*4
	x1 = x - a
	y1 = y - a
	slice_end := [2]int{x1, y1}
	fmt.Println("Square")
	fmt.Println(sq)
	fmt.Println("Perimeter")
	fmt.Println(p)
	fmt.Println("New Coordinates")
	fmt.Println(slice_end)
}
