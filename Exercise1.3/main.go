/*Exercis e 1.3: Experiment to measure the dif ference in running time bet ween our pot ent ial ly
inefficient versions and the one that uses strings.Join . (Section 1.6 illustrates par t of the
time package, and Sec tion 11.4 shows how to write benchmark tests for systematic per-
formance evaluation.)*/

package main

import (
	"fmt"
	"strings"
	"time"
)

func UsingRangedFor(arguments []string) string {
	s, sep := "", ""
	start := time.Now()
	for _, word := range arguments[0:] {
		s += sep + word
		sep = " "
	}
	t := time.Now()
	return "subs:  " + t.Sub(start).String()
}

func UsingJoin(arguments []string) string {
	start := time.Now()
	strings.Join(arguments[0:], " ")
	t := time.Now()
	return "subs:  " + t.Sub(start).String()
}

func main() {
	fmt.Println(UsingRangedFor(arguments))
	fmt.Println(UsingJoin(arguments))
}
