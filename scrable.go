package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	values := map[int]string{
		1:  "aeioulnrst",
		2:  "dg",
		3:  "bcmp",
		4:  "fhvwy",
		5:  "k",
		8:  "jx",
		10: "qz",
	}

	wordFromCommandLine := strings.Join(os.Args[1:], " ")
	word := strings.ToLower(wordFromCommandLine)
	sumOfValues := 0

	for _, wordLetter := range word {
		for key, value := range values {
			for _, char := range value {
				if wordLetter == char {
					sumOfValues += key
				}
			}
		}
	}
	fmt.Println(sumOfValues)
}
