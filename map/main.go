package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red": "#ff0000",
	}

	colors["blue"] = "#00ff00"

	fmt.Println(colors)
}
