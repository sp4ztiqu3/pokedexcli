package main

import (
	//"fmt"
	"strings"
)

func main() {
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}
