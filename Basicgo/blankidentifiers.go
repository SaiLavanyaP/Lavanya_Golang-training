package main

import (
	"fmt"
)

func main() {
	triangle, _, rect := area(2, 10.8, 5.6)
	fmt.Println("area of traingle ,area of rect ", triangle, rect)
}

func area(height, length, breath float32) (float32, float32, float32) {
	var atriangle = (breath * height) / 2
	fmt.Println(atriangle)
	var square = length * length
	var rect = length * breath
	return atriangle, square, rect
}
