package main

import (
	"fmt"
)

func updateThirdElement(p *[4]string) {
	(*p)[3] = "Texas"
}

func Add(a *int, b *int) *int {
	add := *a + *b
	return &add
}
func Sub(a *int, b *int) *int {
	sub := *a - *b
	return &sub
}
func Mul(a *int, b *int) *int {
	mul := *a * *b
	return &mul
}
func Div(a *int, b *int) (*int, *int) {
	q := *a / *b
	r := *a % *b
	return &q, &r
}

func main() {
	//1 creating array with size 5 and printing type of array
	val := [5]int{1, 2, 3, 4, 5}
	fmt.Println("1st Ans:The array values are :", val)
	fmt.Printf("1st Ans:Type of the array:val is %T\n", val)

	//2 creating slice with 10 values assigned and print type of array
	val2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("2nd Ans:values of the slice", val2)
	fmt.Printf("2nd Ans:Type of the slice:val2 is %T\n", val2)

	//3 use slice and create new slice
	val3 := val2[:5]
	val4 := val2[5:]
	val5 := val2[2:7]
	val6 := val2[1:6]
	fmt.Println("3rd Ans:printing slice", val3)
	fmt.Println("3rd Ans:printing slice", val4)
	fmt.Println("3rd Ans:printing slice", val5)
	fmt.Println("3rd Ans:printing slice", val6)

	//4th
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println("4th Ans: Adding 52 to x variable", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("4th Ans: appending slice x and y", x)

	//5th
	data := [4]string{"Missouri", "Kansas", "Chicago", "New york"}
	fmt.Println("5th Ans: Beofore updating data", data)
	updateThirdElement(&data)
	fmt.Println("5th Ans: third element", data)

	//6th Create a calculator app using pointers
	firstval := 100
	secondval := 10
	add := Add(&firstval, &secondval)
	sub := Sub(&firstval, &secondval)
	mul := Mul(&firstval, &secondval)
	q, r := Div(&firstval, &secondval)
	fmt.Println("Addition : ", *add)
	fmt.Println("Subtraction : ", *sub)
	fmt.Println("Multiplication : ", *mul)
	fmt.Printf("Quotient : %v, Remainder : %v ", *q, *r)
}
