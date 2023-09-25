package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Please enter a number : ")
	fmt.Scan(&number)

	if number < 1 {
		fmt.Println("Angka harus lebih besar dari 0")
	}

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	fmt.Println("Even numbers upto your input are :")

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	if number == 1 {
		fmt.Println("Bilangan prima harus lebih atau sama dengan 2")
	}

	// Check is prime
	for x := 2; x <= number; x++ {
		if x == number {
			fmt.Println("The number is a prime number")
			break
		}
		if number%x == 0 {
			fmt.Println("The number is not a prime number")
			break
		}
	}

}
