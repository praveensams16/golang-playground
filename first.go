package main

import (
	"fmt"
)

var i int

func sam(i *int) {
	*i=20
}

func main() {
	i:=10
  sam(&i)
	fmt.Println(i)
}
