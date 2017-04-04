package main

import "fmt"

func main() {
	for index := 1; index <= 100; index++ {
		if index%3 == 0 {
			if index%3 == 0 && index%5 == 0 {
				fmt.Println(index, "FizzBuzz")
			}
			fmt.Println(index, "Fizz")
		} else if index%5 == 0 {
			fmt.Println(index, "Buzz")
			if index%3 == 0 && index%5 == 0 {
				fmt.Println(index, "FizzBuzz")
			}
		} else {
			fmt.Printf("%v \n", index)
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}
