package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])
	symbols := make(map[string]bool)
	for _, char := range string(data) {
		symbol := string(char)
		if len(symbol) == 1 && rune(symbol[0]) < 126 {
			continue //skip ascii
		}
		symbols[symbol] = true
	}
	for key := range symbols {
		fmt.Print(key)
	}
}
