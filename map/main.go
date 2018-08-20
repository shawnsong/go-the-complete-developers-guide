package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string
	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
