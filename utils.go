package gocamp

// For functions I don't know where else to put

import "fmt"

// Print an array of runes
func PrintRuneArray(runes [][]rune) {
	for _, line := range runes {
		for _, e := range line {
			fmt.Printf("%s", string(e))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
