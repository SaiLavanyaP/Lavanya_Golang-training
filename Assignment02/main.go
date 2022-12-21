package main

import (
	"Assignment02/Calculator"
	"fmt"
)

func main() {
	var a, b int
	var operation string
	fmt.Print("Enter first number:")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter Second number:")
	fmt.Scanf("%d", &b)
	fmt.Print("Enter operations(+ - * /)")
	fmt.Scanf("%s", &operation)

	if operation == "+" {
		fmt.Println(Calculator.Add(float32(a), float32(b)))
	} else if operation == "-" {
		fmt.Println(Calculator.Sub(float32(a), float32(b)))
	} else if operation == "*" {
		fmt.Println(Calculator.Multi(float32(a), float32(b)))
	} else if operation == "/" {
		fmt.Println(Calculator.Div(a, b))
	}

}

//Users/sailavanyapeddinti/Desktop/Localrepo/Lavanya_Golang-training
