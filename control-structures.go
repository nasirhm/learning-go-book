package main

import "fmt"

func main(){
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	// Or
	fmt.Println(`1
2
3
4
5
6
7
8
9
10`)

	// For, Go only has one type of loop that can be used in ways you like
	i := 1
	for i <= 10{
		fmt.Println("With for loop :" ,i)
		i = i+1
	}

	// The above program can be written like this as well.
	for i := 1; i<=10; i++{
		fmt.Println("Another way with for loop :", i)
	}
	// If
	if i % 2 == 0{
		println("even")
	} else {
		println("odd")
	}

	// else if
	if i % 2 == 0 {
		println("Divisible by 2")
	}else if i % 3 == 0 {
		println("Divisible by 3")
	}else if i % 4 == 0 {
		println("Divisible by 4")
	}else if i == 11 {
		println("It's 11 and it's a prime")
	}

	// Combining for and if
	for i := 1; i <= 10; i++{
		if i % 2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}

	// Switch Case :
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}

	// Write a program to print out all the numbers evenly divisible by 3 between 1 and 100
	for i := 0; i<100; i++{
		if i % 3 == 0{
			fmt.Println(i)
		}
	}

	// Write a program for Fizz Buzz
	for i:= 0; i<100; i++{
		if i % 3 == 0{
			fmt.Println("Fizz")
		}else if i % 5 ==0{
			fmt.Println("Buzz")
		}else if i % 3 & i%5 ==0{
			fmt.Println("FizzBuzz")
		}else{
			fmt.Println(i)
		}
	}
}