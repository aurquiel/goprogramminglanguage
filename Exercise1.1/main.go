/*Exercis e 1.1: Mo dif y the echo prog ram to als o pr int os.Args[0] , the name of the command
that invoked it.*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func UsingSingleFor(arguments []string) string {
	s, sep := "", ""
	for i := 0; i < len(arguments); i++ {
		s += sep + os.Args[i]
		sep = "\n"
	}
	return "SINGLE: " + s
}

func UsingRangedFor(arguments []string) string {
	s, sep := "", ""
	for _, word := range arguments[0:] {
		s += sep + word
		sep = " "
	}
	return "RANGED: " + s
}

func UsingJoin(arguments []string) string {
	return "JOIN: " + strings.Join(arguments[0:], " ")
}

func main() {
	arguments := os.Args[0:]
	fmt.Println(UsingSingleFor(arguments))
	fmt.Println(UsingRangedFor(arguments))
	fmt.Println(UsingJoin(arguments))
}
