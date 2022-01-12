package main

import "fmt"

func main() {
	// create map variable
	var emptyMap map[string]string
	// construct a map with make func
	constructMap := make(map[string]string)
	// initialize and create map variable
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf756",
	}
	fmt.Println("Empty Map: ", emptyMap)
	fmt.Println("Constructed Map: ", constructMap)
	fmt.Println("Initialized Map: ", colors)
	// add key to map
	colors["white"] = "#000000"
	fmt.Println("After added white key: ", colors)
	// delete key from map
	delete(colors, "white")
	fmt.Println("After deleted white key: ", colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Key: ", k, " - Value: ", v)
	}
}
