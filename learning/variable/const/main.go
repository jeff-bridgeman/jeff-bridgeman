package main

import "fmt"

const p = "death & taxes"

func main() {

	q := 42

	fmt.Println(&q)

	fmt.Println("p - ", p)

	q = q + 1

	fmt.Println("q - ", q)
	fmt.Println("q - ", q)

	a := q

	fmt.Println(a)
	fmt.Println(&a)

}
