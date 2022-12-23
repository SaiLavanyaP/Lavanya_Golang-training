package main

import "fmt"

func main() {
	var num int
	age := 20 //short hand declartion
	fmt.Println(num, age)
	num = 15
	fmt.Println(num)
	age, num = 2, 3
	fmt.Println(age, num)
	fmt.Println(age + num)
	name := "Hi"
	name2 := "hello"
	fmt.Println(name, name2)
	var num2, num3 = 1, 5 //type inference :type can be removed
	fmt.Println(num2, num3)

}

//types of declarations
// variable variablename type = value
//variable variablename =value
// variablename := value
// var a,b,c= 1,1,1 //multiple declarations
