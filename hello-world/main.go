package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var n int
	fmt.Scanln(&n)
	arr := make([]rune, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
}
