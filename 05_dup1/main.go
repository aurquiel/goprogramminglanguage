// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		counts[input.Text()]++
	}
	fmt.Printf("%v", counts)
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
