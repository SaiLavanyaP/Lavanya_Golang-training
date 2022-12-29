package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//1st
	for i := 1; i <= 100; i++ {
		fmt.Printf(" %d", i)
	}
	//2nd
	fmt.Println("Printing odd numbers between 1 to 50")
	for j := 1; j <= 50; j++ {
		if j%2 == 1 {
			fmt.Printf(" %d\n", j)
		}
	}
	//3rd
	fmt.Println("Even numbers are :")
	w := 1
	for {
		if w%2 == 0 {
			fmt.Print(w, " ")
		}
		if w == 50 {
			break
		}
		if w < 50 {
			w = w + 1
		}
	}
	//4th
	fmt.Println("printing quotient")
	var k int
	var num float64
	for k = 50; k <= 105; k++ {
		num = float64(k) / float64(6.0)
		numstring := strconv.FormatFloat(num, 'f', -1, 64)
		fmt.Printf("%v\n", numstring)
	}
	//5th
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the String")
	inp1, _ := reader.ReadString('\n')
	inp1 = strings.TrimSpace(inp1)
	if inp1 == "Golang tutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}
	//6th
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("Golang tutorial")
		} else if i%2 == 0 && i%4 != 0 {
			fmt.Println("Golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		} else {
			fmt.Println(i)
		}
	}
}
