package maths

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Sumdigits HI
func Sumdigits() int {
	//var list []int = os.Open(sums.txt)
	//var list []int

	file, err := os.Open("M:/Jeff Bridgeman/Documents/GitHub/GoWorkspace/src/github.com/jeff.bridgeman/learning/exercises/ProjectEuler/Maths/sums.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list := make([]byte, 6000)
	count, err := file.Read(list)
	if err != nil {
		log.Fatal(err)
	}

	testing, err := ioutil.ReadFile("M:/Jeff Bridgeman/Documents/GitHub/GoWorkspace/src/github.com/jeff.bridgeman/learning/exercises/ProjectEuler/Maths/sums.txt")
	if err != nil {
		log.Fatal(err)
	}

	//lines := strings.Split(string(testing), "\n")
	//var log []int

	fmt.Println("======================")
	fmt.Printf("%q\n", list[:count])
	fmt.Println("Length of list:", len(list))
	fmt.Println("======================")
	fmt.Println(string(testing))
	return count
}
