package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runs(cxt context.Context,c chan string, i *int) {
	for {
	m.Lock()
	*i= *i + 1
	m.Unlock()
	c <- pr(i,"first")
	time.Sleep(1*time.Second)
	select {
	case <-cxt.Done():
		wg.Done()
		return 
	default:
		fmt.Println("from first")
	}
}
}
func pr(i *int,s string) string {
	return fmt.Sprintf("from %s -> %v ",s,*i)
}
func runone(cxt context.Context,c chan string, i *int) {
	for {
	m.Lock()
	*i= *i + 1
	m.Unlock()
	c <- pr(i,"second")
	time.Sleep(1*time.Second)
	select {
	case <-cxt.Done():
		wg.Done()
		return 
	default:
		fmt.Println("from second")
	}
}
}
var (
	wg sync.WaitGroup
	i int
	m sync.Mutex
)

func main() {


	cxt,cancel:=context.WithCancel(context.Background())
	c:=make(chan string , 100)
	wg.Add(1)
	go runs(cxt,c,&i)
	go runone(cxt,c,&i)
	SAM:
	for i:=0;i<200;i++ {
		time.Sleep(2*time.Second)
		fmt.Println( <-c )
		if i == 5 {
			cancel()
			select {
			case <-cxt.Done():
				break  SAM
			default:
				fmt.Println("from main")
			}
		}
	}
	wg.Wait()
}
