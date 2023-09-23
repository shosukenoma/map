package main

import "fmt"

func main() {
	// // Option 1
	// colors := map[string]string {
	// 	"red": "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue": "#0000ff",
	// }

	// // Option 2
	// var colors map[string]string

  // // Option 3
	colors := make(map[string]string)
	colors["white"] = "#fff"

	// Deleting an item by key
	delete(colors, "white")

	fmt.Println(colors)
}