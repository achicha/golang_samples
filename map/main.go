/*learn map DS in golang*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	// declaration
	var colors0 map[string]string

	colors1 := make(map[string]string)
	colors1["white"] = "#ffffff"

	colors2 := map[string]string{
		"red":   "ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	// delete key/value pair
	delete(colors1, "white")

	// iterate over key/value pairs
	printMap(colors2)

	// print out
	fmt.Println(colors0, colors1, colors2)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("color=", key, "hex=", value)
	}
}

func duplicate_count(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}
