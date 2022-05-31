package main

import "fmt"

func main() {
	var a1 = [3]int{1, 2, 3}
	a1[0] = 10
	fmt.Println(a1)

	var a2 = make([]int, 5, 10)
	fmt.Println(a2, len(a2), cap(a2))
}
