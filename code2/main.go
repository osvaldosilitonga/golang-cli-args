package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := os.Args
	if len(length) < 3 {
		fmt.Println("Invalid args. Please insert args <operation> <operand1> <operand2>")
		return
	}

	operation := os.Args[1]
	operand1 := os.Args[2]
	val1, _ := strconv.ParseInt(operand1, 10, 64)
	operand2 := os.Args[3]
	val2, _ := strconv.ParseInt(operand2, 10, 64)

	switch operation {
	case "add":
		fmt.Println("Result :", val1+val2)
	case "sub":
		fmt.Println("Result :", val1-val2)
	case "mul":
		fmt.Println("Result :", val1*val2)
	case "div":
		fmt.Println("Result :", val1/val2)
	default:
		fmt.Println("Operation not found!")
	}
}
