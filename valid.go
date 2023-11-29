package main

import "fmt"

type Sam struct {
	i int
}

func (i *Sam)run() bool {
	if i < 10 {
		i=i+1
		return true
	} else {
		return false
	}
	
}

func main() {
     
	s:=Sam{i:0}
	for i:=0;s.run();{
		fmt.Println("hiii")
	}
}