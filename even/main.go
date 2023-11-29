package main

import (
	"fmt"
	"local/find"
)

func main() {
	d:=find.Sam{}
	
	for i:=0;i<100;i++ {
		if find.Odd(i) {
			fmt.Println("even number",i)
		} else {
			fmt.Println("odd number",i)
		}
		fmt.Println("Finding the number",d.Add(i,2))
	}
}
