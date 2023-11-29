package main

import (
	"flag"
	"fmt"
)
type Sam struct {
	i int
}

type Sam1 struct {
	i int
}

func (i Sam)run(j int) bool {
	if  j % i.i == 0 {
		i.i=i.i+1
		return true
	} else {
		return false
	}
	
}

func (i Sam1)run(j int) bool {
	if  j % i.i == 0 {
		i.i=i.i+1
		return true
	} else {
		return false
	}
	
}

type Inter interface {
	run(int) bool
}

func main() {
	d:=flag.Int("count",10,"an int")
     
	var I Inter = Sam{i:8}
	var I1 Inter = Sam1{i:7}

	
	for i:=0;i<*d;i++ {
		if I.run(i) {
			fmt.Println(i)
		} else if I1.run(i) {
			fmt.Println("bye",i)
		}
	}
}