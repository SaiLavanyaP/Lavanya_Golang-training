package main

import (
	"fmt"
)

func main() {
	a, b := 1.2, 1.3
	fmt.Println("sum", a+b)
	//complex type: real and imaginary type 6+7i
	valu := 6 + 7i
	fmt.Println(valu)
	//implicity declaration
	var message = "hi" //here data we are passing can be anytype, so type of varibale is assumed by the operands
	fmt.Println(message)
	//explicity
	var mess string = "Hello" //here mess can be of string type only
	fmt.Println(mess)
	//Type Conversion
	var i = 55
	j := 67.5
	fmt.Println(i + int(j))
	var s = 5.1
	var k int = int(s)
	fmt.Println(k)
	//string conversion uses value := strconv.Itoa(i)
	var w string = "good morning"
	fmt.Printf("%s sweety.\n", w)
}
