package main

import "fmt"

func main() {
	//map[key]value{k1:v1,}

	//colors := make(map[string]string)
	// var colors map[string]string
	colors := map[string]string{
		"red":   "0xFF0000",
		"blue":  "0x00FF00",
		"green": "0x0000FF",
	}
	colors["white"] = "0xFFFFFF"

	//delete(colors, "red")
	fmt.Println(colors)

	fmt.Println("\nIterating over map:")
	for color, hex := range colors {
		fmt.Println(color + " has hexcode " + hex)
	}
}
