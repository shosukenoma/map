package main

import "fmt"

func main() {
	// Option 1
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	// // Option 2
	// var colors map[string]string

  // // Option 3
	// colors := make(map[string]string)

	// Adding an item to a map
	// colors["white"] = "#fff"

	// // Deleting an item by key
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Maps are reference types
// So the reference data is copied, not the map itself.