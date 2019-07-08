package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: dldist string1 string2")
		fmt.Println("  The program will compute the Damerau-Levenshtein distance between string1 and string2.")
		return
	}
	dist := DLDist(os.Args[1], os.Args[2])
	fmt.Println(dist)
}
