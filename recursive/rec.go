package main

import "fmt"

func adds(n int) int {
	var y int
	if len(string(y)) == 0 {
		return y
	} else {
		y = n + adds(n-1)
	}
	return y
}

func main() {
	d:=adds(10)
	fmt.Println(d)
}
