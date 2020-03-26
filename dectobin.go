package main

import "fmt"

var n, input int

func main() {
	fmt.Printf("Please enter the decimal number \n")

	fmt.Scanln(&input)

	n = 1

	for n*2 <= input {
		n = n * 2
	}

	for n >= 1 {
		fmt.Print(input / n)
		input = input % n
		n = n / 2
	}
}
