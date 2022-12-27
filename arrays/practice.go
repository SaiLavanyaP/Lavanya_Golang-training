package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)
	b := []int{1, 2, 3}
	fmt.Println(b)
	c := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length is", len(c))
}
