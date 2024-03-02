package main

import "fmt"

func main() {
	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	fmt.Println(colors)
	fmt.Println(colors["red"])
}
