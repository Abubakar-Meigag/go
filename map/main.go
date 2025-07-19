package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
		"black": "#000000",
	}

	// Adding elements in a map
	colors["yellow"] = "#ffff00"
	// deleting an element from a map
	delete(colors, "white")
	// updating an element in a map
	colors["red"] = "#f00"

	fmt.Println(colors)
	printMap(colors)
	fmt.Println(colors["red"])
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}