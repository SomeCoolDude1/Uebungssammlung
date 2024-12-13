package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	for i := Lower; i <= Upper; i++ {
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println("Fizz")
			}
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println()

}
