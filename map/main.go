package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"black": "000000",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
