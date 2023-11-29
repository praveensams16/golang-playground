package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"
)

type Sam struct {
	a, b int
	c    chan int
}

type Runner struct {
	a string `json:"a"`
	b string `json:"b"`
	h string `json:"h"`
}

func (s *Sam) run(j int, wg *sync.WaitGroup) {
	s.c <- s.a + s.b + j
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	x := 0
	c := make(chan int, 1000)
	d := Sam{a: 1, b: 2, c: c}
	de, e := ioutil.ReadFile("sam.go")
	if e == nil {
		defer fmt.Println(strings.Split(string(de), "\n")[4])
	}
bro:

	for {
		x += 1
		if x == 100 {
			break bro
		}
		wg.Add(1)
		go d.run(x, &wg)
	}
	wg.Wait()
complete:
	for {

		select {
		case d := <-c:
			fmt.Println(d)
		default:
			fmt.Println("completed")
			break complete
		}

	}
	r := "1.3"
	fmt.Println(reflect.TypeOf(r))
	g := Runner{a: "a", b: "b", h: "run"}
	fmt.Println(g.a)
	fmt.Scan(&g.a)
	fmt.Println(g.a)
}
