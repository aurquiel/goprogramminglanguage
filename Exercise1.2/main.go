/*Exercis e 1.2: Mo dif y the echo prog ram to print the index and value of each
of its arguments, on e per line.*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func UsingrangedFor(arguments []string) string {
	s := ""
	for index, word := range arguments[0:] {
		s += strconv.Itoa(index) + ", " + word + "\n"
	}
	return s
}

func main() {
	arguments := os.Args[0:]
	fmt.Println(UsingrangedFor(arguments))
}
