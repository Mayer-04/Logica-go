package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 10)
	slice = append(slice, 20)
	slice = append(slice, 30)
	fmt.Println(slice)
}
