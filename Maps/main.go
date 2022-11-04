package main

import "fmt"

func main() {
	//map[key]value{k1:v1,}
	colors := map[string]string{
		"red":   "0xFF0000",
		"blue":  "0x00FF00",
		"green": "0x0000FF",
	}
	fmt.Println(colors)
}
