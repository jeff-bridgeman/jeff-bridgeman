package maths

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//Sumdigits HI
func Sumdigits() int {

	file, err := ioutil.ReadFile("M:/Jeff Bridgeman/Documents/GitHub/GoWorkspace/src/github.com/jeff.bridgeman/learning/exercises/ProjectEuler/Maths/sums.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	var nums []string
	nums = make([]string, len(lines))
	for _, n := range lines { //Credit to Mustafa on Stack Overflow
		//a, err := strconv.Atoi(n)
		//fmt.Println(n)
		nums = append(nums, n)
		if err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println("======================")
	fmt.Println(lines)
	return 42
}
