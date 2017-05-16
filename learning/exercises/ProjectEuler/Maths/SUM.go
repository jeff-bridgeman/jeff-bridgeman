package maths

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Sumdigits HI
func Sumdigits() float64 {

	file, err := ioutil.ReadFile("M:/Jeff Bridgeman/Documents/GitHub/GoWorkspace/src/github.com/jeff.bridgeman/learning/exercises/ProjectEuler/Maths/sums.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\r\n")
	var nums []float64
	nums = make([]float64, len(lines))

	fmt.Println("After creating nums: ", len(nums))

	var sums float64
	sums = 0

	for _, n := range lines { //Credit to Mustafa on Stack Overflow
		//fmt.Println(n)
		a, err := strconv.ParseFloat(n, 64)
		//fmt.Println(a)
		sums = sums + a
		nums = append(nums, a)
		if err != nil {
			log.Fatal(err)
		}
	}

	return sums
}
