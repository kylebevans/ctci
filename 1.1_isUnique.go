package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isUnique(s string) bool {
	var unique bool = true
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				unique = false
			}
		}
	}
	return unique
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string")

	inString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: ", err)
	}

	result := isUnique(inString)
	if result {
		fmt.Println("All characters in your string are unique")
	} else {
		fmt.Println("All characters in your string are not unique")
	}
	return
}
